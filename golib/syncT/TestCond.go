package syncT

import (
	"fmt"
	"sync"
	"time"
)

var locker1 = new(sync.Mutex)
var cond1 = sync.NewCond(locker1)

func test(x int) {
	cond1.L.Lock() //获取锁
	cond1.Wait()   //等待通知  暂时阻塞
	fmt.Println(x)
	time.Sleep(time.Second * 1)
	cond1.L.Unlock() //释放锁
}
func main() {
	for i := 0; i < 40; i++ {
		go test(i)
	}
	fmt.Println("start all")
	time.Sleep(time.Second * 3)
	fmt.Println("broadcast")
	cond1.Signal() // 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 3)
	cond1.Signal() // 3秒之后 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 3)
	cond1.Broadcast() //3秒之后 下发广播给所有等待的goroutine
	time.Sleep(time.Second * 60)
}
