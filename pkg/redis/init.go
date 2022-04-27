package redis

import (
	"fmt"
	"github.com/tiptok/gocomm/pkg/cache"
	"github.com/tiptok/gocomm/pkg/cache/gzcache"
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
	// 默认redis实例
	//cache.InitDefault(
	//	cache.WithDefaultRedisPool(redis.GetRedisPool()),
	//	cache.WithDebugLog(true, log.Logger),
	//)

	// 单个redis实例
	//cache.InitMultiLevelCache(cache.WithDebugLog(true, log.Logger)).
	//	RegisterCache(gzcache.NewNodeCache(redisSource, constant.REDIS_AUTH))

	// 多个redis实例
	cache.InitMultiLevelCache(cache.WithDebugLog(true, log.Logger)).
		RegisterCache(gzcache.NewClusterCache([]string{redisSource}, constant.REDIS_AUTH)) //"127.0.0.1:6379", "127.0.0.1:6380", "127.0.0.1:6381"
}
