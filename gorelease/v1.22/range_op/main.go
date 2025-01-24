package main

import (
	"fmt"
	"sync"
)

func main() {
	//rangeAsync()
	//rangeNumber()
	rangeIterator()

}

func rangeAsync() {
	fmt.Println("1.rangeAsync")
	group := sync.WaitGroup{}
	list := []string{"A", "B", "C", "D", "E", "F"}
	for i, s := range list {
		group.Add(1)
		go func() {
			defer group.Done()
			fmt.Println(i, s)
		}()
	}
	group.Wait()
}

func rangeNumber() {
	fmt.Println("2.rangeNumber: range 的范围前闭后开 [0, N)。")
	for i := range 10 {
		fmt.Println(i)
	}
}

func rangeIterator() {
	s := Set[int]{m: make(map[int]struct{})}
	s.m[1] = struct{}{}
	s.m[2] = struct{}{}
	s.m[4] = struct{}{}
	s.Iterator(func(v int) bool {
		//fmt.Println(v)
		return true
	})

	for e := range s.Iterator {
		fmt.Println(e)
	}
}

type Set[E comparable] struct {
	m map[E]struct{}
}

func (s *Set[E]) Iterator(f func(E) bool) {
	for e := range s.m {
		if !f(e) {
			return
		}
	}
}
