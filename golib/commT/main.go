package main

import (
	"log"
	"runtime"
)

func main() {
	defer func() {
		if p := recover(); p != nil {
			log.Print(p)
			var buf [2 << 10]byte
			log.Println(string(buf[:runtime.Stack(buf[:], true)]))
		}
	}()
	slice := make([]string, 2, 4)
	example(slice, "hello", 10)
}

func example(slice []string, str string, i int) {
	panic("Want stack trace")
}
