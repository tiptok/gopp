package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

type SseHandler struct {
	clients map[chan string]struct{}
}

func NewSseHandler() *SseHandler {
	return &SseHandler{
		clients: make(map[chan string]struct{}),
	}
}

/*
SSE（Server-Sent Events）是 HTML5 提供的一种技术，允许服务器通过持久化的 HTTP 连接向客户端单向推送事件。
相比 WebSocket，SSE 更轻量，支持简单的实时更新场景，且基于标准 HTTP 协议，开箱即用
doc: https://mp.weixin.qq.com/s/tf5Opz2I8VCu5n24YjXUDw
访问: http://127.0.0.1:8080/static/
*/

// Serve 处理 SSE 连接
func (h *SseHandler) Serve(w http.ResponseWriter, r *http.Request) {
	// 设置 SSE 必需的 HTTP 头
	w.Header().Add("Content-Type", "text/event-stream")
	w.Header().Add("Cache-Control", "no-cache")
	w.Header().Add("Connection", "keep-alive")

	// 为每个客户端创建一个 channel
	clientChan := make(chan string)
	h.clients[clientChan] = struct{}{}

	// 客户端断开时清理
	defer func() {
		delete(h.clients, clientChan)
		close(clientChan)
	}()

	// 持续监听并推送事件
	for {
		select {
		case msg := <-clientChan:
			// 发送事件数据
			fmt.Fprintf(w, "data: %s\n\n", msg)
			w.(http.Flusher).Flush()
		case <-r.Context().Done():
			// 客户端断开连接
			return
		}
	}
}

// SimulateEvents 模拟周期性事件
func (h *SseHandler) SimulateEvents() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		message := fmt.Sprintf("Server time: %s", time.Now().Format(time.RFC3339))
		// 广播给所有客户端
		for clientChan := range h.clients {
			select {
			case clientChan <- message:
			default:
				// 跳过阻塞的 channel
			}
		}
	}
}

func main() {
	// 创建 go-zero REST 服务，集成静态文件服务
	server := rest.MustNewServer(rest.RestConf{
		Host: "0.0.0.0",
		Port: 8080,
	}, rest.WithFileServer("/static", http.Dir("static")))
	defer server.Stop()

	// 初始化 SSE 处理
	sseHandler := NewSseHandler()

	// 注册 SSE 路由
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/sse",
		Handler: sseHandler.Serve,
	}, rest.WithTimeout(0))

	// 在单独的 goroutine 中模拟事件
	go sseHandler.SimulateEvents()

	logx.Info("Server starting on :8080")
	server.Start()
}
