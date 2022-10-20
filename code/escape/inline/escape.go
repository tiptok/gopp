package main

import "fmt"

func main() {
	// fmt.Println("Called stackAnalysis", stackAnalysis())

	fmt.Println("Called InlineStackAnalysis", InlineStackAnalysis())

	//fmt.Println("Called ptrStackAnalysis", ptrStackAnalysis())
}

//go:noinline
func stackAnalysis() int {
	data := 100
	return data
}

//go:inline
func InlineStackAnalysis() int {
	data := 100
	return data
}

//go:noinline
func ptrStackAnalysis() *int {
	data := 100
	return &data
}

/*  stackAnalysis
# go build -gcflags=-m escape.go

.\escape.go:6:13: inlining call to fmt.Println
.\escape.go:6:13: ... argument does not escape
.\escape.go:6:14: "Called stackAnalysis" escapes to heap
.\escape.go:6:51: stackAnalysis() escapes to heap
*/

/*  InlineStackAnalysis
# go build -gcflags=-m escape.go

.\escape.go:20:6: can inline InlineStackAnalysis
.\escape.go:8:63: inlining call to InlineStackAnalysis
.\escape.go:8:13: inlining call to fmt.Println
.\escape.go:8:13: ... argument does not escape
.\escape.go:8:14: "Called InlineStackAnalysis" escapes to heap
.\escape.go:8:63: ~R0 escapes to heap
.\escape.go:27:2: moved to heap: data

*/

/* ptrStackAnalysis
# go build -gcflags=-m escape.go

.\escape.go:10:13: inlining call to fmt.Println
.\escape.go:20:6: can inline InlineStackAnalysis
.\escape.go:27:2: moved to heap: data
.\escape.go:10:13: ... argument does not escape
.\escape.go:10:14: "Called ptrStackAnalysis" escapes to heap

*/
