package syncT

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestRunScenario(t *testing.T) {
	RunScenario()
}

func RunScenario() {
	numCpu := runtime.NumCPU()
	taskSize := 10
	jobs := make(chan int, taskSize)
	results := make(chan int, taskSize)
	var wg sync.WaitGroup

	workExampleFunc := func(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		for job := range jobs {
			res := job * 2
			fmt.Printf("Work %d do things,produce result %d\n", id, res)
			time.Sleep(time.Millisecond * time.Duration(100))
			results <- res
		}
	}

	for i := 0; i < numCpu; i++ {
		wg.Add(1)
		go workExampleFunc(i, jobs, results, &wg)
	}

	totalTasks := 100

	go func() {
		defer wg.Done()
		for i := 0; i < totalTasks; i++ {
			n := <-results
			fmt.Printf("Get results %d \n", n)
		}
		close(results)
	}()

	for i := 0; i < totalTasks; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

func TestTicker(t *testing.T) {
	tm := time.NewTicker(time.Second)
	index := 0
	for {
		select {
		case <-tm.C:
			index++
			fmt.Printf("数量:%d", index)
			if index > 100 {
				tm.Stop()
				return
			}
		}
	}
}
