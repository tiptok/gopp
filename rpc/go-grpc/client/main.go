package main

import (
	"context"
	"flag"
	"github.com/tiptok/gocomm/common"
	"github.com/tiptok/gopp/pkg/protobuf/user"
	"google.golang.org/grpc"
	"log"
)

var serverAddr string

func init() {
	flag.StringVar(&serverAddr, "addr", "127.0.0.1:8080", "rpc server address")
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	client := user.NewUserClient(conn)
	rsp, _ := client.GetUser(context.Background(), &user.GetUsersReq{Id: 1})
	log.Println(common.JsonAssertString(rsp))
}
