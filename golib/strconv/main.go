package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("strconv delete")
	dst := make([]byte, 0)
	print(strconv.AppendBool(dst, false))
}

func print(args ...interface{}) {
	fmt.Println(args...)
}
