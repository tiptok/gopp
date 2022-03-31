package gozero

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tiptok/gocomm/sync/task"
	"github.com/zeromicro/go-zero/core/syncx"
)

// 用途：防止缓存穿透，在并发情况下多次访问数据库，加载缓存，只允许一个线程访问数据库，其他线程共享结果
func Test_Example_SharedCalls(t *testing.T) {
	var count int32
	sharedCalls := syncx.NewSingleFlight()
	c := make(chan string)
	getValue := func() (interface{}, error) {
		atomic.AddInt32(&count, 1)
		time.Sleep(time.Second * 10)
		return <-c, nil
	}
	var num = 10
	gt := task.NewGroupTask()
	gt.WithWorkerNumber(num)

	for i := 0; i < num; i++ {
		gt.Run(func() {
			v, ex, e := sharedCalls.DoEx("test", getValue)
			if e != nil {
				t.Fatal(e)
			}
			fmt.Println(v, ex)
		})
	}
	time.Sleep(100 * time.Millisecond)
	c <- "foo"
	gt.Wait()

	if count != 1 {
		fmt.Printf("want =%v got=%v ", 1, count)
	}
}

// 测试：wait 可以在多个协程并发存在
func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	for i := 0; i < 10; i++ {
		temp := i
		go func(v int) {
			wg.Wait()
			fmt.Printf("wait done %v\n", temp)
		}(temp)
	}
	time.Sleep(time.Second * 5)
	wg.Done()
}

// 用途：只允许一个线程执行，其他的要等得到锁以后才能继续
func Test_LockedCalls(t *testing.T) {
	lockedCalls := syncx.NewLockedCalls()

	var count int32
	getValue := func() (interface{}, error) {
		atomic.AddInt32(&count, 1)
		time.Sleep(time.Millisecond * 100)
		return nil, nil
	}
	var num = 3
	gt := task.NewGroupTask()
	gt.WithWorkerNumber(num)

	for i := 0; i < num; i++ {
		gt.Run(func() {
			_, e := lockedCalls.Do("test", getValue)
			if e != nil {
				t.Fatal(e)
			}
		})
	}
	gt.Wait()
	assert.EqualValues(t, num, count)
}

func TestAtomicBool(t *testing.T) {
	b := syncx.NewAtomicBool()
	b.Set(true)
	if b.True() {
		b.CompareAndSwap(true, false)
	}
	assert.Equal(t, b.True(), false)
}

func TestAtomicFloat64(t *testing.T) {
	f := syncx.NewAtomicFloat64()
	f.Set(1.5)
	f.Add(6.5)
	assert.EqualValues(t, f.Load(), 8)
	f.CompareAndSwap(8, 6)
	assert.EqualValues(t, f.Load(), 6)

	v := math.Float32bits(1.5)
	fv := math.Float32frombits(v)
	assert.EqualValues(t, fv, 1.5)
}

// 用途:资源释放时候，标记已经关闭 Close() 关闭信号，确保退出
func TestDoneChan(t *testing.T) {
	var wg sync.WaitGroup
	quit := syncx.NewDoneChan()
	timeTimer := time.NewTimer(time.Second)
	wg.Add(1)
	go func() {
		select {
		case <-quit.Done():
			wg.Done()
		case <-timeTimer.C:
			wg.Done()
			fmt.Println("time out")
		}
	}()
	time.Sleep(time.Millisecond * 900)
	quit.Close()
	quit.Close() //只会关闭一次
	wg.Wait()
}

// 用途：对象池  限制数量，限制存活时间
func TestPool(t *testing.T) {
	limit := 10
	var value int32
	stack := syncx.NewPool(limit, func() interface{} {
		return atomic.AddInt32(&value, 1)
	}, func(interface{}) {

	})

	tg := task.NewGroupTask()
	tg.WithWorkerNumber(100)
	for i := 0; i < 100; i++ {
		tg.Run(func() {
			v := stack.Get().(int32)
			if v <= 0 {

			}
			time.Sleep(time.Millisecond * 50)
			stack.Put(v)
		})
	}
	time.Sleep(time.Millisecond * 500)
	tg.Wait()
	assert.EqualValues(t, 10, value)
}
