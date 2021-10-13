package main

import "testing"

func Test_Trace(t *testing.T) {
	slice := make([]string, 2, 4)
	example(slice, "hello", 10)
}
