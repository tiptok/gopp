package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/tiptok/gocomm/common"
	"github.com/tiptok/gocomm/pkg/log"
	"github.com/tiptok/gopp/pkg/protobuf/user"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

func main() {
	flag.Parse()

	client := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"localhost:2379"},
			Key:   "zrpc",
		},
	})
	conn := client.Conn()
	userClient := user.NewUserClient(conn)

	resp, err := userClient.GetUser(context.Background(), &user.GetUsersReq{
		Id: 1,
	})
	if err == nil {
		log.Info(fmt.Sprintf("get user :%v", common.AssertString(resp)))
	}
}
