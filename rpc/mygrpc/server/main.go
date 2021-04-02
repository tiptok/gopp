package main

import (
	"flag"
	"github.com/tiptok/gopp/pkg/protobuf/user"
	_ "github.com/tiptok/gopp/pkg/redis"
	"github.com/tiptok/gopp/rpc/mygrpc/balancer"
	"github.com/tiptok/gopp/rpc/pzrpc/server/controller"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	svcName  = "gopp/userAdmin"
	addr     = "127.0.0.1:8080"
	etcdAddr = "127.0.0.1:2379"
)

/*
go run rpc/mygrpc/server/main.go -addr 127.0.0.1:8081
go run rpc/mygrpc/server/main.go -addr 127.0.0.1:8082
go run rpc/mygrpc/server/main.go -addr 127.0.0.1:8083

go run rpc/mygrpc/client/main.go
*/
func main() {
	flag.StringVar(&addr, "addr", addr, "addr to lis")
	flag.Parse()

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	} else {
		log.Println("listen:", addr)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()
	user.RegisterUserServer(grpcServer, controller.NewUserController(addr))
	reflection.Register(grpcServer)

	// 服务健康检查
	grpc_health_v1.RegisterHealthServer(grpcServer, controller.NewHealthController())

	go balancer.Register(etcdAddr, svcName, addr, 5)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)

	go func() {
		s := <-ch
		balancer.UnRegister(svcName, addr)
		if i, ok := s.(syscall.Signal); ok {
			os.Exit(int(i))
		} else {
			os.Exit(0)
		}
	}()

	grpcServer.Serve(lis)
}
