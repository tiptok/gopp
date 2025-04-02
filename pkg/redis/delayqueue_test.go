package redis

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"testing"
	"time"
)

func TestRedisDelayQueue(t *testing.T) {
	// 初始化 Redis 客户端
	ctx := context.Background()
	client, _ := redis.NewRedis(redis.RedisConf{Host: "localhost:6379", Pass: "", Type: "node"})
	// 消费者函数示例
	consumerFunc := func(ctx context.Context, message string) error {
		t.Log("consumer msg ==>", message, time.Now().Unix())

		//time.Sleep(1 * time.Second) // 模拟处理耗时
		if message == "msg-5" {
			return NewDelayRetryError(nil, time.Second*5)
		} else if message == "msg-6" {
			return RetryErr
		}
		return nil
	}

	// 初始化延迟队列
	queue := NewDelayQueue(ctx,
		client,
		"gopp:{delay-queue}:test",
		consumerFunc,
		3,
		time.Second*10)

	// 开始消费消息
	queue.Start()

	for i := 1; i <= 10; i++ {
		err := queue.addMessage(fmt.Sprintf("msg-%d", i), 5*time.Second)
		t.Log("producer msg ==>", fmt.Sprintf("msg-%d", i), time.Now().Unix())
		if err != nil {
			t.Fatalf("add message error: %s", err)
		}
	}

	// 模拟程序运行一段时间
	time.Sleep(20 * time.Second)

	// 关闭延迟队列
	queue.Shutdown()
}
