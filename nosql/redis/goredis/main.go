package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {

}

func ExampleRedis() {
	// Redis client
	//rdb:=redis.NewClient(&redis.Options{
	//	Addr: "localhost:6379",
	//	Password: "",
	//	DB: 0,
	//})

	// Redis Sentinel
	//redis.NewFailoverClient()
	// 允许把只读命令随机路由到redis 从节点
	rdb := redis.NewFailoverClusterClient(&redis.FailoverOptions{
		MasterName:     "mymaster",
		SentinelAddrs:  []string{":26379", ":26378", ":26377"},
		RouteByLatency: true,
		RouteRandomly:  true,
	})
	rdb.AddHook(CustomHook{})
	ctx := context.Background()
	val := rdb.Do(ctx, "get", "test")
	if val.Err() != nil {
		panic(val.Err())
	}
	fmt.Println(val.String())

	statusCmd := rdb.Do(ctx, "set", "test", "ok3")
	if statusCmd.Err() != nil {
		fmt.Println(statusCmd.Err())
		return
	}

	fmt.Println(statusCmd.Result())
}

// 自定义钩子，统计，链路追踪等
type CustomHook struct {
}

func (hook CustomHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	fmt.Println("before ", cmd.Name(), cmd.String())
	return ctx, nil
}
func (hook CustomHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	return nil
}
func (hook CustomHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	return ctx, nil
}
func (hook CustomHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	return nil
}
