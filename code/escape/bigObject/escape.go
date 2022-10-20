package main

// 栈空间不足而导致的逃逸(空间开辟过大)

func InitSlice() {
	s := make([]int, 1000, 1000) // not escape
	s = make([]int, 1000, 10000) // escape

	for index := range s {
		s[index] = index
	}
}

func VarInitSlice() {
	l := 20
	c := make([]int, 0, l) // 堆 动态分配不定空间 逃逸
	_ = c
}

func main() {
	InitSlice()
	VarInitSlice()
}

/*
$ go build -gcflags=-m escape.go

# command-line-arguments
.\escape.go:3:6: can inline InitSlice
.\escape.go:12:6: can inline main
.\escape.go:13:11: inlining call to InitSlice
.\escape.go:4:11: make([]int, 1000, 1000) does not escape
.\escape.go:5:10: make([]int, 1000, 10000) escapes to heap
.\escape.go:13:11: make([]int, 1000, 1000) does not escape
.\escape.go:13:11: make([]int, 1000, 10000) escapes to heap


当栈空间不足以存放当前对象,或无法判断当前切片长度时,会将对象分配到堆中
*/
