package routers

import (
	"github.com/gin-gonic/gin/ginS"
	"github.com/tiptok/gocomm/pkg/myrest"
	"github.com/tiptok/gocomm/pkg/myrest/handler"
	"github.com/tiptok/gopp/http/gin/controllers"
	"github.com/tiptok/gopp/pkg/constant"
)

func init() {
	ginS.Use(myrest.GinMiddlewareChains(
		handler.TracingHandler,
		handler.RecoverHandler(),
		handler.LogHandler, //handler.DetailedLogHandler
		handler.LimitConnHandler(constant.MaxConn),
		//handler.TimeoutHandler(constant.TimeOutDuration),   //请求改变了。使用有问题
		handler.LimitBytesHandler(constant.MaxSize),
	))
}

func init() {
	c := &controllers.UserController{}

	ginS.POST("/v1/user/", c.CreateUser)
	ginS.PUT("/v1/user/:userId", c.UpdateUser)
	ginS.GET("/v1/user/:userId", c.GetUser)
	ginS.DELETE("/v1/user/:userId", c.DeleteUser)
	ginS.GET("/v1/user/", c.ListUser)
}
