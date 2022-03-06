package testT

import (
	"log"
	"testing"
	"time"
)

//go test  -run TestWillTimeOut -timeout 1s
//go test time out
func TestWillTimeOut(t *testing.T) {
	log.Println("start")
	time.Sleep(time.Second * 2)
	log.Println("end.")
}

//go test -run TestShortFlag -short
//go test short mode flag:长时间运行的测试会被跳过
func TestShortFlag(t *testing.T) {
	if testing.Short() {
		log.Println("skip short flag")
		t.Skip("skip in short mode")
	}
	log.Println("short flag")
}

//go test -run TestVerboseDebug -v
//go test verbose flag ：打印详细的额外日志
func TestVerboseDebug(t *testing.T) {
	if testing.Verbose() {
		log.Println("verbose debug info")
	}
	log.Println("simple info")
}

//GOMAXPROCS=2 go test to test with 2 CPUs
//go test -cpu=1,2,4  运行三次    ex:go test -cpu=1,2,3 -run TestVerboseDebug
//go test -parallel n  并行测试   ex:go test -parallel 2 -run TestParallet

func TestParallet(t *testing.T) {
	t.Parallel()
}

func BenchmarkParallet(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Println("simple Benchmark")
		}
	})
}
