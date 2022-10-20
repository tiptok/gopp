package main

type Student struct {
	Name string
	Age  int
}

func StudentRegister(name string, age int) *Student {
	s := new(Student) //局部变量s逃逸到堆

	s.Name = name
	s.Age = age

	return s
}

func main() {
	StudentRegister("dashen", 18)
}

/*
$ go build -gcflags=-m escape.go

# command-line-arguments
.\escape.go:8:6: can inline StudentRegister
.\escape.go:17:6: can inline main
.\escape.go:18:17: inlining call to StudentRegister
.\escape.go:8:22: leaking param: name
.\escape.go:9:10: new(Student) escapes to heap
.\escape.go:18:17: new(Student) does not escape

s 虽然为 函数StudentRegister()内的局部变量, 其值通过函数返回值返回. 但s 本身为指针类型. 所以其指向的内存地址不会是栈而是堆.
*/
