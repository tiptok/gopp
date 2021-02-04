package main

import (
	"github.com/astaxie/beego"
	_ "github.com/tiptok/gopp/http/beego/routers"
)

func main() {
	beego.BConfig.CopyRequestBody = true

	beego.Run(":8080")
}
