package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"github.com/tiptok/gopp/pkg/constant"
	_ "github.com/tiptok/gopp/pkg/constant"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	"github.com/tiptok/gocomm/pkg/log"
)

var listenAddr string
var remoteServeAddr string
var serviceName string

func flagParse() {
	flag.StringVar(&serviceName, "name", "server", "服务名称")
	flag.StringVar(&listenAddr, "laddr", "localhost:10081", "服务监听地址")
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

	{
		// set up a span reporter
		reporter := zipkinhttp.NewReporter(constant.TRACE_REPORTER_URL)
		defer reporter.Close()

		// create our local service endpoint
		endpoint, err := zipkin.NewEndpoint(serviceName, listenAddr)
		if err != nil {
			log.Error("unable to create local endpoint: %+v\n", err)
		}

		// initialize our tracer
		nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
		if err != nil {
			log.Error("unable to create tracer: %+v\n", err)
		}

		// use zipkin-go-opentracing to wrap our tracer
		tracer := zipkinot.Wrap(nativeTracer)

		// optionally set as Global OpenTracing tracer instance
		opentracing.SetGlobalTracer(tracer)
	}

	http.Handle("/echo", OpenTracingAdapter(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if len(remoteServeAddr) == 0 {
			time.Sleep(time.Millisecond * (time.Duration(rand.Int63n(int64(500)))))
			writer.Write([]byte("host:" + listenAddr))
			return
		}
		req, _ := http.NewRequest(http.MethodGet, "http://"+remoteServeAddr+"/echo", nil)
		// 传递 tranceId spanId 到下一个服务
		span := opentracing.SpanFromContext(request.Context())
		if span != nil {
			opentracing.GlobalTracer().Inject(
				span.Context(),
				opentracing.HTTPHeaders,
				opentracing.HTTPHeadersCarrier(req.Header))
		}
		searchInfo(request.Context())
		rsp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Error(err)
		}
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

// searchInfo 自定义函数传入链路信息
func searchInfo(ctx context.Context) {
	if rand.Intn(2) == 1 {
		span, _ := opentracing.StartSpanFromContext(ctx, "searchInfo")
		time.Sleep(time.Millisecond * (time.Duration(rand.Int63n(int64(500)))))
		defer span.Finish()
	}
}

// OpenTracingAdapter 中间件 拦截或发起一个trace
func OpenTracingAdapter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var sp opentracing.Span
		opName := r.URL.String()
		// Attempt to join a trace by getting trace context from the headers.
		wireContext, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header))
		var ctx context.Context
		if err != nil {
			// If for whatever reason we can't join, go ahead an start a new root span.
			//sp = opentracing.StartSpan(opName)
			sp, ctx = opentracing.StartSpanFromContext(r.Context(), opName)
		} else {
			//sp = opentracing.StartSpan(opName, opentracing.ChildOf(wireContext))
			sp, ctx = opentracing.StartSpanFromContext(r.Context(), opName, opentracing.ChildOf(wireContext))
		}
		defer sp.Finish()
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
