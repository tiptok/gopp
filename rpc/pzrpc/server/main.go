package main

import (
	"flag"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"github.com/tiptok/gopp/pkg/protobuf/user"
	_ "github.com/tiptok/gopp/pkg/redis"
	"github.com/tiptok/gopp/rpc/pzrpc/server/controller"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//var configFile = flag.String("f", "etc/config.json", "the config file")
var configFile = flag.String("f", "etc/pzrpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c zrpc.RpcServerConf
	conf.MustLoad(*configFile, &c)

	server := zrpc.MustNewServer(c, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, controller.NewUserController())
		reflection.Register(grpcServer)
	})

	//interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	//	st := time.Now()
	//	resp, err = handler(ctx, req)
	//	log.Info(fmt.Sprintf("method: %s time: %v\n", info.FullMethod, time.Since(st)))
	//	return resp, err
	//}
	//
	//server.AddUnaryInterceptors(interceptor)
	server.Start()
}
