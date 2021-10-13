package main

import (
	"flag"
	"fmt"
)

var (
	port    = flag.Int("Port", 808, "监听端口")
	address = flag.String("Adr", "127.0.0.1", "监听地址")
)

//go run main.go -Port 8000 -Adr "192.168.3.54" -User "tiptok"
func main() {
	var user string
	flag.StringVar(&user, "User", "admin", "登录用户")
	fmt.Println(flag.Args())
	flag.Parse()
	fmt.Println(*address, *port, user)
}
