package routers

import (
	"github.com/tiptok/gopp/http/gozero/controllers"
)

func init() {
	c := &controllers.UserController{}

	ServerRouter.POST("/v1/user/", c.CreateUser)
	ServerRouter.PUT("/v1/user/:userId", c.UpdateUser)
	ServerRouter.GET("/v1/user/:userId", c.GetUser)
	ServerRouter.DELETE("/v1/user/:userId", c.DeleteUser)
	ServerRouter.GET("/v1/user/", c.ListUser)
}
