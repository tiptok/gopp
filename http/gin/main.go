package main

import (
	"github.com/gin-gonic/gin/ginS"
	_ "github.com/tiptok/gopp/http/gin/routers"
	_ "github.com/tiptok/gopp/pkg/redis"
)

func main() {
	ginS.Run(":8080")
}
