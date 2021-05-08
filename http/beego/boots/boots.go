package boots

import (
	"fmt"
	"github.com/tiptok/gocomm/pkg/cache/gzcache"
	"github.com/tiptok/gocomm/pkg/mybeego"
	"github.com/tiptok/gocomm/pkg/myrest/handler"
	"github.com/tiptok/gopp/pkg/constant"
)

func init() {
	nodeCache := gzcache.NewNodeCache(fmt.Sprintf("%v:%v", constant.REDIS_HOST, constant.REDIS_PORT), "")

	// 需要再main函数优先init全局的拦截器
	mybeego.Use(handler.TracingHandler,
		handler.RecoverHandler(),
		handler.LogHandler, //handler.DetailedLogHandler
		handler.LimitConnHandler(constant.MaxConn),
		handler.TimeoutHandler(constant.TimeOutDuration),
		handler.LimitBytesHandler(constant.MaxSize),
		handler.AtomicPersistenceQueryHandler(
			handler.WithRequestQueryHashFunc(handler.ComputeHttpRequestQueryHash),
			handler.WithServiceName(constant.ServiceName),
			handler.WithCache(nodeCache),
			handler.WithExpire(360),
			handler.WithRouters([]string{
				"/user/{userId}",
				"/user",
			}),
		),
	)
}
