package main

import (
	server "github.com/gin-gonic/gin/ginS"
	"github.com/tiptok/gocomm/pkg/myrest"
	"github.com/tiptok/gocomm/pkg/myrest/handler"
	_ "github.com/tiptok/gopp/http/gin/routers"
	"github.com/tiptok/gopp/pkg/constant"
	_ "github.com/tiptok/gopp/pkg/redis"
)

func init() {
	server.Use(myrest.GinMiddlewareChains(
		handler.TracingHandler,
		handler.RecoverHandler(),
		handler.LogHandler, //handler.DetailedLogHandler
		handler.LimitConnHandler(constant.MaxConn),
		//handler.TimeoutHandler(constant.TimeOutDuration),   //请求改变了。使用有问题
		handler.LimitBytesHandler(constant.MaxSize),
	))
}

func main() {
	server.Run(":8080")
}
