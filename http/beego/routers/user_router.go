package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/tiptok/gocomm/pkg/myrest"
	"github.com/tiptok/gocomm/pkg/myrest/handler"
	"github.com/tiptok/gopp/http/beego/controllers"
	"github.com/tiptok/gopp/pkg/constant"
)

func init() {
	c := &controllers.UserController{}
	beego.Post("/v1/user/", bindMiddleware(c.CreateUser))
	beego.Put("/v1/user/:userId", bindMiddleware(c.UpdateUser))
	beego.Get("/v1/user/:userId", bindMiddleware(c.GetUser))
	beego.Delete("/v1/user/:userId", bindMiddleware(c.DeleteUser))
	beego.Get("/v1/user/", bindMiddleware(c.ListUser))
}

func bindMiddleware(work func(c *context.Context)) func(c *context.Context) {
	return myrest.BeeUseMiddleware(work,
		handler.TracingHandler,
		handler.RecoverHandler(),
		handler.LogHandler, //handler.DetailedLogHandler
		handler.LimitConnHandler(constant.MaxConn),
		handler.TimeoutHandler(constant.TimeOutDuration),
		handler.LimitBytesHandler(constant.MaxSize),
	)
}
