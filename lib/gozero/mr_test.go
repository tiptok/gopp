package gozero

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/mr"
	"log"
	"testing"
	"time"
)

func Test_mr(t *testing.T) {
	val, err := mr.MapReduce(func(source chan<- interface{}) {
		// generator
		for i := 0; i < 10; i++ {
			source <- i
		}
	}, func(item interface{}, writer mr.Writer, cancel func(error)) {
		// mapper
		i := item.(int)
		time.Sleep(time.Second)
		writer.Write(Query{
			Index: i,
		})
	}, func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
		// reducer
		var data = make(map[int]int, 0)
		for i := range pipe {
			item := i.(Query)
			data[item.Index] = item.Index
		}
		writer.Write(data)
	}, mr.WithWorkers(5))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result:", val)
}

type Query struct {
	Index int
	Data  struct{}
}
