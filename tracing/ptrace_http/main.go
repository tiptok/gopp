package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/tiptok/gopp/pkg/constant"
	_ "github.com/tiptok/gopp/pkg/constant"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/tiptok/gocomm/pkg/log"
	"github.com/tiptok/gocomm/pkg/myrest/handler"
	"github.com/tiptok/gocomm/pkg/trace"
)

var listenAddr string
var remoteServeAddr string
var serviceName string

func flagParse() {
	flag.StringVar(&serviceName, "name", "server", "服务名称")
	flag.StringVar(&listenAddr, "laddr", "localhost:10082", "服务监听地址")
	flag.StringVar(&remoteServeAddr, "rmaddr", "", "远程服务地址")
	flag.Parse()

	log.Debug(fmt.Sprintf("serviceName = %v listenAddr=%v remoteServeAddr=%v", serviceName, listenAddr, remoteServeAddr))
}

/*
	启动三个服务
	HTTP GET  localhost:10081/echo

	go run main.go -name sv1 -laddr localhost:10081 -rmaddr localhost:10082
	go run main.go -name sv2 -laddr localhost:10082 -rmaddr localhost:10083
	go run main.go -name sv3 -laddr localhost:10083

*/

func main() {
	flagParse()
	trace.BindEndpoint(serviceName, listenAddr)
	trace.BindZipkinReporter(constant.TRACE_REPORTER_URL)

	http.Handle("/echo", handler.TracingHandler(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if len(remoteServeAddr) == 0 {
			time.Sleep(time.Millisecond * (time.Duration(rand.Int63n(int64(500)))))
			writer.Write([]byte("host:" + listenAddr))
			return
		}
		req, _ := http.NewRequest(http.MethodGet, "http://"+remoteServeAddr+"/echo", nil)
		// 传递 tranceId spanId 到下一个服务
		trace.HttpInject(request.Context(), req)
		rsp, _ := http.DefaultClient.Do(req)
		defer func() {
			if rsp != nil {
				rsp.Body.Close()
			}
		}()
		data, _ := io.ReadAll(rsp.Body)
		buf := bytes.NewBuffer([]byte(serviceName + " host:" + listenAddr))
		buf.WriteString("\n")
		buf.Write(data)
		writer.Write(buf.Bytes())
	})))
	http.ListenAndServe(listenAddr, nil)
}
