package main

import (
	"context"
	"flag"
	"time"

	"github.com/cloudwego/netpoll"
	"github.com/tiptok/gocomm/pkg/log"
)

var network = "tcp"
var address = "127.0.0.1:8888"
var s = 5

func main() {
	flag.StringVar(&address, "remote", address, "远程服务地址")
	log.Debug(network, address)

	// 创建连接
	conn, err := netpoll.DialConnection(network, address, 5000*time.Millisecond)
	if err != nil {
		panic("dial netpoll connection fail")
	}

	// 设置读事件回调
	conn.SetOnRequest(newOnRequest())

	for {
		// conn write & flush message
		conn.Writer().WriteBinary([]byte("hello world"))
		conn.Writer().Flush()
		time.Sleep(time.Second * time.Duration(s))
	}
}

func newOnRequest() netpoll.OnRequest {
	return func(ctx context.Context, connection netpoll.Connection) error {
		log.Info("callbak -> onRequest",connection.RemoteAddr(),connection.LocalAddr())
		return nil
	}
}
