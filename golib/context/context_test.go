package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Test_Context_Cancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go func(c context.Context) {
		for {
			select {
			case <-c.Done():
				fmt.Println("work exit with cancel")
				return
			default:
				fmt.Println("do work ", time.Now().Unix())
				time.Sleep(time.Second)
			}
		}
	}(ctx)

	time.Sleep(time.Second * 5)
	cancel()
	time.Sleep(time.Second * 2)
	fmt.Println("exit (0)")
}

func Test_Context_Deadline(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*2)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overselpt")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
}
