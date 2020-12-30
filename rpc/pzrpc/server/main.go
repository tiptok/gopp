package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"github.com/tiptok/gocomm/pkg/log"
	"github.com/tiptok/gopp/rpc/pzrpc/protocol"
	"github.com/tiptok/gopp/rpc/pzrpc/server/handler"
	"google.golang.org/grpc"
	"time"
)

var configFile = flag.String("f", "etc/config.json", "the config file")

func main() {
	flag.Parse()

	var c zrpc.RpcServerConf
	conf.MustLoad(*configFile, &c)

	server := zrpc.MustNewServer(c, func(grpcServer *grpc.Server) {
		protocol.RegisterUserServer(grpcServer, handler.NewUserServer())
	})
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		st := time.Now()
		resp, err = handler(ctx, req)
		log.Info(fmt.Sprintf("method: %s time: %v\n", info.FullMethod, time.Since(st)))
		return resp, err
	}

	server.AddUnaryInterceptors(interceptor)
	server.Start()
}
