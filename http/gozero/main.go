package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tiptok/gocomm/pkg/log"
	"github.com/tiptok/gocomm/pkg/mygozero"
	_ "github.com/tiptok/gopp/http/gozero/routers"
	_ "github.com/tiptok/gopp/pkg/redis"
)

var configFile = flag.String("f", "conf/api.yaml", "the config file")

func init() {

}

func main() {
	var restConf rest.RestConf
	flag.Parse()
	conf.MustLoad(*configFile, &restConf)

	log.InitGzLog(logx.LogConf{ServiceName: "gopp", Mode: "file"}) //,Path: "logs"

	server := rest.MustNewServer(restConf)
	server.AddRoutes(mygozero.ServerRouter.Routers)
	defer server.Stop()

	log.Info(fmt.Sprintf("Starting server at %s:%d...\n", restConf.Host, restConf.Port))
	server.Start()
}
