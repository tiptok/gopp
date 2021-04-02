package main

import (
	"context"
	"github.com/tiptok/gocomm/common"
	"github.com/tiptok/gopp/pkg/constant"
	"github.com/tiptok/gopp/pkg/protobuf/user"
	"github.com/tiptok/gopp/rpc/mygrpc/balancer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	_ "google.golang.org/grpc/health" //健康检查
	"google.golang.org/grpc/resolver"
	"log"
	"math/rand"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	builder := balancer.NewResolver(constant.ETCD_ADDRESS)
	resolver.Register(builder)

	conn, err := grpc.Dial(builder.Scheme()+"://gopp/userAdmin", grpc.WithBalancerName(roundrobin.Name), grpc.WithInsecure())
	//conn,err :=grpc.Dial(builder.Scheme() +"/gopp/userAdmin", grpc.WithBalancerName(roundrobin.Name), grpc.WithInsecure())
	//conn,err :=grpc.Dial("zrpc/7587853421409661042", grpc.WithBalancerName(roundrobin.Name), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := user.NewUserClient(conn)

	for {
		resp, err := client.GetUser(context.Background(), &user.GetUsersReq{
			Id: int64(rand.Intn(10)),
		})
		if err != nil {
			log.Println("get error response :", err)
		} else {
			log.Println("get response from:", resp.Host, "data", common.JsonAssertString(resp))
		}
		<-time.After(time.Second * 3)
	}

}
