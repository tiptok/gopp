package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strings"
	"time"

	//"github.com/redis/go-redis/v9"
	"golang.org/x/sync/semaphore"

	task "github.com/tiptok/gopp/pkg/common"
)

const (
	CreateQueueName = "{confirmation:create}"
	DeleteQueueName = "{confirmation:delete}"
)

var (
	RetryErr      = errors.New("err retry")
	RetryNowError = errors.New("err retry right now")
)

type DelayQueue struct {
	client       *redis.Redis
	ctx          context.Context
	cancel       context.CancelFunc
	queueName    string
	ackQueueName string
	consumerFunc func(context.Context, string) error
	weighted     *semaphore.Weighted
	done         chan struct{}
	retryTime    time.Duration // err retry time
	*task.Sleeper
}

// FormatQueueNameInCluster 集群下需要对队列名做下格式化
func FormatQueueNameInCluster(queueName string) string {
	if strings.Contains(queueName, "{") && strings.Contains(queueName, "}") {
		return queueName
	}
	// 处理集群情况Lua脚本问题,queue打上标签键值都落在同一个节点上
	if !strings.HasPrefix(queueName, "{") {
		queueName = "{" + queueName
	}
	if !strings.HasSuffix(queueName, "}") {
		queueName += "}"
	}
	return queueName
}

// NewDelayQueue 初始化延迟队列
func NewDelayQueue(ctx context.Context, client *redis.Redis, queueName string, consumerFunc func(context.Context, string) error, weightedSize int64, retryTime time.Duration) *DelayQueue {
	ctx, cancel := context.WithCancel(ctx)
	return &DelayQueue{
		ctx:          ctx,
		cancel:       cancel,
		client:       client,
		queueName:    queueName,
		ackQueueName: queueName + ":ack",
		consumerFunc: consumerFunc,
		weighted:     semaphore.NewWeighted(weightedSize),
		retryTime:    retryTime,
		done:         make(chan struct{}),
	}
}

// Start 开始消费消息
func (dq *DelayQueue) Start() {
	dq.Sleeper = task.NewDefaultSleeper(dq.ctx.Done())
	//dq.reprocessUnAckMessages()
	go dq.consumeMessages()
}

// 消费 Redis 中的消息并处理
func (dq *DelayQueue) consumeMessages() {
	defer func() {
		dq.done <- struct{}{}
	}()

	for {
		select {
		case <-dq.ctx.Done():
			return
		default:
			// 从延迟队列中获取消息
			messages := dq.getMessages()
			if messages == nil || len(messages) == 0 {
				// 如果没有消息，等待一段时间后重试
				dq.Sleep()
				continue
			}
			// 处理每一条消息
			for _, message := range messages {
				msg, ok := message.(string)
				if !ok {
					continue
				}
				err := dq.weighted.Acquire(dq.ctx, 1)
				if err != nil {
					break
				}

				go func(msg string) {
					defer dq.weighted.Release(1)
					e := dq.consumerFunc(dq.ctx, msg)
					if e != nil {
						// 如果不是重试错误，记录错误日志
						if errors.Is(e, RetryErr) {
							logx.Errorf("consumerFunc msg [%v] err: %v", msg, e)
						}
						// 根据不同的错误类型进行不同的处理
						var addErr error
						var delayErr = &DelayRetryError{}
						switch {
						case errors.Is(e, RetryNowError):
							// 立即重试
							addErr = dq.addMessage(msg, 0)
						case errors.As(e, delayErr):
							// 延迟重试
							addErr = dq.addMessage(msg, delayErr.DelayTime)
						default:
							// 按照默认重试时间重试
							addErr = dq.addMessage(msg, dq.retryTime)
						}
						// 检查 addMessage 函数的返回值
						if addErr != nil {
							logx.Errorf("Failed to add message [%v] for retry: %v", msg, addErr)
						}
						return
					}
					// 消费成功后确认消息
					dq.ackMessage(msg)
				}(msg)
			}

			dq.Reset()
		}
	}
}

// 获取消息
func (dq *DelayQueue) getMessages() []any {
	// Lua 脚本：获取消息、从延迟队列删除并添加到 ack 队列
	const luaScript = `
    local items = redis.call('ZRANGEBYSCORE', KEYS[1], '-inf', ARGV[1], 'LIMIT', 0, ARGV[2])
    if #items == 0 then
        return nil  -- 如果没有符合条件的消息，返回空
    end
    
    for i, message in ipairs(items) do
        -- 从延迟队列删除消息
        redis.call('ZREM', KEYS[1], message)
        -- 将消息添加到 ack 队列
        redis.call('ZADD', KEYS[2], ARGV[3], message)
    end
    return items
`
	result, err := dq.client.EvalCtx(dq.ctx, luaScript,
		[]string{dq.queueName, dq.ackQueueName},
		time.Now().Unix(), 10, time.Now().Unix(),
	)
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			logx.Errorf("getMessages err: %v", err)
		}
		return nil
	}

	items, ok := result.([]any)
	if !ok {
		logx.Errorf("getMessages result type err [%v]", result)
		return nil
	}

	return items
}

// 处理完成删除
func (dq *DelayQueue) ackMessage(message any) {
	_, err := dq.client.ZremCtx(dq.ctx, dq.ackQueueName, message)
	if err != nil {
		logx.Errorf("ackMessage ZRem [%v] err: %v", message, err)
	}
}

// 重新处理未确认的消息
func (dq *DelayQueue) reprocessUnAckMessages() {
	// Lua 脚本：重新处理未确认的消息
	const reprocessLuaScript = `
    local unackMessages = redis.call('ZRANGEBYSCORE', KEYS[1], '-inf', ARGV[1])
    if #unackMessages == 0 then
        return nil -- 如果没有符合条件的未确认消息，返回空
    end

    for _, message in ipairs(unackMessages) do
        -- 从 ack 队列中移除消息
        redis.call('ZREM', KEYS[1], message)
        -- 将消息重新放回延迟队列，使用当前时间作为分数
        redis.call('ZADD', KEYS[2], ARGV[2], message)
    end
    return unackMessages
`
	now := time.Now().Unix()
	// 调用 Lua 脚本处理未确认的消息
	_, err := dq.client.EvalCtx(
		dq.ctx,
		reprocessLuaScript,
		[]string{dq.ackQueueName, dq.queueName},
		now, // 当前时间，用于筛选超时的消息
		now, // 当前时间，用作新的分数
	)

	if err != nil && !errors.Is(err, redis.Nil) {
		logx.Errorf("reprocessUnAckMessages err: %v", err)
		return
	}
}

// Shutdown 关闭
func (dq *DelayQueue) Shutdown() error {
	dq.cancel()
	<-dq.done
	fmt.Println(fmt.Sprintf("--DelayQueue %s Shutdown", dq.queueName))
	return nil
}

// 添加消息到延迟队列
func (dq *DelayQueue) addMessage(message string, delay time.Duration) error {
	_, err := dq.client.ZaddCtx(dq.ctx, dq.queueName, time.Now().Add(delay).Unix(), message)
	if err != nil {
		logx.Errorf("addMessage msg [%v], err: %v", message, err)
		return err
	}
	return nil
}

// DelayRetryError 延迟重试错误
type DelayRetryError struct {
	Err       error
	DelayTime time.Duration
}

func (e DelayRetryError) Error() string {
	var sb strings.Builder
	sb.WriteString("retry with delay:%s")
	args := []interface{}{e.DelayTime}

	if e.Err != nil {
		sb.WriteString(", err:%v")
		args = append(args, e.Err)
	}
	return fmt.Sprintf(sb.String(), args...)
}

func NewDelayRetryError(err error, delay time.Duration) DelayRetryError {
	return DelayRetryError{Err: err, DelayTime: delay}
}
