package context

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func Test_Server_Studown(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
		<-sigchan
		cancel()
	}()

	svr := &Server{}
	svr.Run(ctx) // ← long running process
	log.Println("graceful stop")
}

type Server struct{}

func (s *Server) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("cancel received, attempting graceful stop...")
			// clean up process
			return
		default:
			handleRequest()
		}
	}
}

func handleRequest() {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
}
