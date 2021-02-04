package main

import (
	"github.com/astaxie/beego"
	_ "github.com/tiptok/gopp/http/beego/routers"
	_ "github.com/tiptok/gopp/pkg/redis"
)

func main() {
	beego.BConfig.CopyRequestBody = true

	beego.Run(":8080")
}
