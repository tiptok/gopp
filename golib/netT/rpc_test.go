package netT

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"testing"
	"time"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func TestRpcInvoke(t *testing.T) {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
	time.Sleep(time.Second * 2)
	go clientInvoke(t)
	time.Sleep(time.Second * 2)
}
func clientInvoke(t *testing.T) {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := &Args{
		10, 10,
	}
	var reply int
	if err = client.Call("Arith.Multiply", args, &reply); err != nil {
		t.Fatal(err)
	}
	t.Log(fmt.Printf("Arith.Multiply(): %d*%d=%d", args.A, args.B, reply))

	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	_ = <-divCall.Done
	t.Log(fmt.Printf("Arith.Divide(): %d/%d=%d", args.A, args.B, (*quotient).Quo))
}
