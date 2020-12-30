package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/discov"
	"github.com/tal-tech/go-zero/zrpc"
	"github.com/tiptok/gocomm/common"
	"github.com/tiptok/gocomm/pkg/log"
	"github.com/tiptok/gopp/rpc/pzrpc/protocol"
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
	userClient := protocol.NewUserClient(conn)

	resp, err := userClient.GetUser(context.Background(), &protocol.GetUsersReq{
		Id: 1,
	})
	if err == nil {
		log.Info(fmt.Sprintf("get user :", common.AssertString(resp)))
	}
}
