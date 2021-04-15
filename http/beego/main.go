package main

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/tiptok/gocomm/pkg/mybeego"
	"github.com/tiptok/gocomm/pkg/myrest/handler"
	_ "github.com/tiptok/gopp/http/beego/routers"
	"github.com/tiptok/gopp/pkg/constant"
	_ "github.com/tiptok/gopp/pkg/redis"
)

func init() {
	// 需要再main函数优先init全局的拦截器
	mybeego.Use(handler.TracingHandler,
		handler.RecoverHandler(),
		handler.LogHandler, //handler.DetailedLogHandler
		handler.LimitConnHandler(constant.MaxConn),
		handler.TimeoutHandler(constant.TimeOutDuration),
		handler.LimitBytesHandler(constant.MaxSize))
}

func main() {
	beego.BConfig.CopyRequestBody = true
	beego.Run(":8080")
}
