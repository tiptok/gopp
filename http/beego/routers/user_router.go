package routers

import (
	server "github.com/tiptok/gocomm/pkg/mybeego"
	"github.com/tiptok/gopp/http/beego/controllers"
)

func init() {
	c := &controllers.UserController{}
	server.POST("/v1/user/", c.CreateUser)
	server.PUT("/v1/user/:userId", c.UpdateUser)
	server.GET("/v1/user/:userId", c.GetUser)
	server.DELETE("/v1/user/:userId", c.DeleteUser)
	server.GET("/v1/user/", c.ListUser)
}
