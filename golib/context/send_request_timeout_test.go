package context

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestSendRequestTimeOut(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	SendRequest(ctx)
}

func SendRequest(ctx context.Context) {
	respCh := make(chan interface{}, 1)
	go sendRequest(respCh)

	select {
	case <-ctx.Done():
		log.Println("operation timed out!")
	case <-respCh:
		log.Println("response received")
	}
}

func sendRequest(ch chan<- interface{}) {
	time.Sleep(60 * time.Second)
	ch <- struct{}{}
}
