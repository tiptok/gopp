package main

import (
	"context"
	"github.com/tiptok/gopp/pkg/constant"
	_ "github.com/tiptok/gopp/pkg/constant"
	"time"

	"github.com/tiptok/gocomm/pkg/trace"
)

// tracing zipkin
// zipkin quickly start https://zipkin.io/pages/quickstart.html
// TODO:研究zipkin httpReport上报流程

func main() {
	trace.BindEndpoint("ptrace", "localhost")
	trace.BindZipkinReporter(constant.TRACE_REPORTER_URL)
	go func() {
		ctx := context.Background()
		ctx, span := trace.StartServerSpan(ctx, nil, "ptrace", "start serve")
		defer span.Finish()
		login(ctx)
		getUserInfo(ctx)
	}()

	time.Sleep(time.Second * 30)
}

/*
	trace 执行流程
    main -> login
            -> findUser
               -> mysqlUserRepository
         -> getUserInfo
*/

func login(ctx context.Context) {
	ctx, span := trace.StartClientSpan(ctx, "ptrace", "login")
	defer span.Finish()
	time.Sleep(time.Second * 1)
	findUser(ctx)
}

func findUser(ctx context.Context) {
	ctx, span := trace.StartClientSpan(ctx, "ptrace", "findUser")
	defer span.Finish()
	time.Sleep(time.Second * 1)
	mysqlUserRepository(ctx)
}

func mysqlUserRepository(ctx context.Context) {
	_, span := trace.StartClientSpan(ctx, "ptrace", "mysqlUserRepository")
	defer span.Finish()
	time.Sleep(time.Second * 1)
}

func getUserInfo(ctx context.Context) {
	_, span := trace.StartClientSpan(ctx, "ptrace", "getUserInfo")
	defer span.Finish()
	time.Sleep(time.Second * 1)
}
