package redis

import (
	"fmt"
	"github.com/tiptok/gocomm/pkg/cache"
	"github.com/tiptok/gocomm/pkg/log"
	"github.com/tiptok/gocomm/pkg/redis"
	"github.com/tiptok/gopp/pkg/constant"
)

func init() {
	redisSource := fmt.Sprintf("%v:%v", constant.REDIS_HOST, constant.REDIS_PORT)
	err := redis.InitWithDb(100, redisSource, constant.REDIS_AUTH, "0")
	if err != nil {
		log.Error(err)
	}
	cache.InitDefault(
		cache.WithDefaultRedisPool(redis.GetRedisPool()),
		cache.WithDebugLog(true, log.DefaultLog),
	)
}
