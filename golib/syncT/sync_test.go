package syncT

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

/*
	Cond
*/
var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

func TestCond(t *testing.T) {
	runtime.GOMAXPROCS(2)
	for i := 10; i < 100; i++ {
		go NotifyCond(i)
	}
	t.Log("end")
	fmt.Println("Begin Notify:")
	cond.Signal()
	time.Sleep(time.Second * 3)
	cond.Signal()
	time.Sleep(time.Second * 3)
	cond.Signal()
	time.Sleep(time.Second * 3)
	fmt.Println("Begin Broadcast:")
	cond.Broadcast()
	time.Sleep(time.Second * 20)
	fmt.Println("End")
}

func NotifyCond(val int) {
	cond.L.Lock()
	cond.Wait()
	fmt.Printf("Value:%2d", val)
	time.Sleep(time.Second * 1)
	cond.L.Unlock()
}

/*sync.Once */

func TestOnceDo(t *testing.T) {
	ch := make(chan int, 4)
	once := new(sync.Once)
	for i := 0; i < 4; i++ {
		go func(x int) {
			once.Do(func() {
				fmt.Println("Once Do :", x)
			})
			fmt.Println("Values:", x)
			ch <- 1
		}(i)
	}
	for i := 0; i < 4; i++ {
		<-ch
	}
	t.Log("End")
}

/*sync.WaitGroup*/
func TestWaitGroup(t *testing.T) {
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go ProcessWG(wg, i)
	}
	wg.Wait()
	fmt.Println("Wait Group Done")
}

func ProcessWG(wg *sync.WaitGroup, i int) {
	fmt.Printf("process :%d is on\n", i)
	// if i != 9 {
	// 	wg.Done()
	// }
	wg.Done() //阻塞直到所有的
}

func TestTimer(t *testing.T) {
	tick := time.NewTicker(time.Second * 1)
	go func() {
		for t := range tick.C {
			fmt.Println("tick at", t)
		}
	}()
	time.Sleep(time.Second * 100)
}
