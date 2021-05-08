package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/tiptok/gopp/http/beego/boots"
	_ "github.com/tiptok/gopp/http/beego/routers"
	_ "github.com/tiptok/gopp/pkg/redis"
)

func main() {
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.Listen.EnableAdmin = true
	beego.Run(":8080")
}
