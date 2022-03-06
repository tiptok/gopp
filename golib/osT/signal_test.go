package osT

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func Test_Signal(t *testing.T) {
	log.Println("signal begin")
	{
		osSignal := make(chan os.Signal, 1)
		signal.Notify(osSignal, os.Interrupt, os.Kill, syscall.SIGTERM)
		<-osSignal
	}
	log.Println("signal end")
}

func Test_handler(t *testing.T) {
	h := new(handler)
	h.set(1)
	if h.want(1) {
		log.Println("equal :", 1)
	}
	if h.want(4) {
		log.Println("equal :", 4)
	}
	h.set(4)
	if h.want(4) {
		log.Println("equal :", 4)
	}
	h.set(60)
}

const (
	numSig = 65 // max across all systems
)

//存放状态
type handler struct {
	mask [(numSig + 31) / 32]uint32
}

func (h *handler) want(sig int) bool {
	return (h.mask[sig/32]>>uint(sig&31))&1 != 0
}

func (h *handler) set(sig int) {
	h.mask[sig/32] |= 1 << uint(sig&31)
}

func (h *handler) clear(sig int) {
	h.mask[sig/32] &^= 1 << uint(sig&31)
}
