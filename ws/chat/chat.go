package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var BootName = "@chatai"
var clients []websocket.Conn

var listen string

func main() {
	flag.StringVar(&listen, "l", ":38080", "监听地址")
	flag.Parse()
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		clients = append(clients, *conn)

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err.Error())
				//return
			}

			// Print the message to the console
			fmt.Printf("%s 接收: %s\n", conn.RemoteAddr(), string(msg))
			if strings.HasPrefix(string(msg), BootName) {
				data := strings.TrimLeft(string(msg), BootName)
				response, err := Completions(data)
				if err != nil {
					fmt.Println(err.Error())
					response = "发生错误:" + err.Error()
				}
				conn.WriteMessage(msgType, []byte(response))
				continue
			}
			for _, client := range clients {
				// Write message back to browser
				if err = client.WriteMessage(msgType, msg); err != nil {
					return
				}
			}

		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.ListenAndServe(listen, nil)
}
