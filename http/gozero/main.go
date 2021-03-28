package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tiptok/gopp/http/gozero/routers"
	_ "github.com/tiptok/gopp/pkg/redis"
)

var configFile = flag.String("f", "conf/api.yaml", "the config file")

func main() {
	var restConf rest.RestConf
	flag.Parse()
	conf.MustLoad(*configFile, &restConf)
	server := rest.MustNewServer(restConf)
	server.AddRoutes(routers.ServerRouter.Routers)
	defer server.Stop()

	fmt.Printf("Starting server at %s:%d...\n", restConf.Host, restConf.Port)
	server.Start()
}
