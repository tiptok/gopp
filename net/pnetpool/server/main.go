package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/netpoll"
	"github.com/tiptok/gocomm/pkg/log"
)

var network = "tcp"
var address = "127.0.0.1:8888"

func main() {
	// 创建 listener
	listener, err := netpoll.CreateListener(network, address)
	if err != nil {
		panic("create netpoll listener fail")
	}

	// handle: 连接读数据和处理逻辑
	var onRequest netpoll.OnRequest = handler
	var onPrepare netpoll.OnPrepare = prepare

	var opts = []netpoll.Option{
		netpoll.WithIdleTimeout(1 * time.Second),
		// handle: 连接读数据和处理逻辑
		netpoll.WithIdleTimeout(10 * time.Minute),
		netpoll.WithOnPrepare(onPrepare),
	}
	eventLoop, err := netpoll.NewEventLoop(onRequest, opts...)
	if err != nil {
		panic("create netpoll event-loop fail")
	}
	// 运行 Server
	err = eventLoop.Serve(listener)
	if err != nil {
		panic("netpoll server exit")
	}
}

// 读事件处理
func handler(ctx context.Context, connection netpoll.Connection) error {
	reader := connection.Reader()
	bts, err := reader.Next(reader.Len())
	if err != nil {
		return err
	}
	log.Debug(fmt.Sprintf("Key: %s, data: %s", ctx.Value(defaultKey), string(bts)))

	connection.Write([]byte("echo:" + string(bts)))
	return connection.Writer().Flush()
}

type defaultKeyType struct{}
var defaultKey defaultKeyType = defaultKeyType{}

func prepare(connection netpoll.Connection) context.Context {
	return context.WithValue(context.Background(), defaultKey, "context")
}
