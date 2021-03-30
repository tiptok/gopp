package main

import (
	"flag"
	"fmt"
	"github.com/tiptok/gopp/pkg/protobuf/user"
	_ "github.com/tiptok/gopp/pkg/redis"
	"github.com/tiptok/gopp/rpc/pzrpc/server/controller"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var port int

func init() {
	flag.IntVar(&port, "p", 8080, "listen port")
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	user.RegisterUserServer(grpcServer, controller.NewUserController())
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}
