package main

import "fmt"

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := Fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Printf("Fibonacci: %d\n", f())
	}
}

/*
$ go build -gcflags=-m escape.go
# command-line-arguments
.\escape.go:5:6: can inline Fibonacci
.\escape.go:7:9: can inline Fibonacci.func1
.\escape.go:14:16: inlining call to Fibonacci
.\escape.go:7:9: can inline main.func1
.\escape.go:17:34: inlining call to main.func1
.\escape.go:17:13: inlining call to fmt.Printf
.\escape.go:6:2: moved to heap: a
.\escape.go:6:5: moved to heap: b
.\escape.go:7:9: func literal escapes to heap
.\escape.go:14:16: func literal does not escape
.\escape.go:17:13: ... argument does not escape
.\escape.go:17:34: ~R0 escapes to heap

Fibonacci()函数中原本属于局部变量的a和b,由于闭包的引用,不得不将二者放到堆上,从而产生逃逸
*/

/*
总结

1.逃逸分析在编译阶段完成
2.逃逸分析目的是决定内分配地址是栈还是堆
3.栈上分配内存比在堆中分配内存有更高的效率
4.栈上分配的内存不需要GC处理
5.堆上分配的内存使用完毕会交给GC处理

通过逃逸分析,不逃逸的对象分配在栈上,当函数返回时就回收了资源,不需gc标记清除,从而减少gc的压力
同时,栈的分配比堆快,性能好(逃逸的局部变量会在堆上分配,而没有发生逃逸的则有编译器在栈上分配
*/
