package netT

import (
	"log"
	"net"
	"testing"
	"time"
)

func TestIPParse(t *testing.T) {
	ip := net.ParseIP("192.168.3.87")
	ipB := net.ParseIP("192.168.3.86")
	log.Printf("IP:%v\n", ip)
	log.Println(ip.IsGlobalUnicast())
	log.Println("DefaultMask:", ip.DefaultMask())
	log.Println("IP Equal:", ip.Equal(ipB))
	log.Println("IP Mask:", ip.Mask(net.IPv4Mask(192, 168, 0, 0)))
}

func TestIPNet(t *testing.T) {
	ip, ipnet, err := net.ParseCIDR("192.168.100.1/31")
	if err != nil {
		t.Fatal(err)
	}
	log.Println("IP:", ip)
	log.Println("IPNet:", ipnet)
	log.Println("IPNet.Network:", ipnet.Network())
	log.Println("IPNet.String:", ipnet.String())

	/*
		type Addr interface {
		    Network() string // 网络名
		    String() string  // 字符串格式的地址
		}
	*/

}

func TestConn(t *testing.T) {
	c, err := net.Dial("tcp", "127.0.0.1:9927")
	if err != nil {
		log.Println("Dial Error:", err)
	}
	c.SetWriteDeadline(time.Now().Add(time.Second * 5))
	for {
		_, err := c.Write([]byte("Hello World"))
		if err != nil {
			log.Println(c.RemoteAddr(), " Close Conn", err)
		}
		time.Sleep(time.Second * 10)
	}
}
