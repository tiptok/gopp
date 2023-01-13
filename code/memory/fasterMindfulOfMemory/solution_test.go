package fasterMindfulOfMemory

import (
	"bytes"
	"io"
	"log"
	"os"
	"testing"
)

func Example() {
	file, err := os.Open("testdata/input.txt")
	if err != nil {
		log.Fatalf("could not open input file: %v", err)
	}
	defer file.Close()
	if err := Solve(file, os.Stdout); err != nil {
		log.Fatalf("could not solve: %v", err)
	}
}

func BenchmarkNoMemoryAlloc(b *testing.B) {
	input, err := os.ReadFile("testdata/input.txt")
	if err != nil {
		b.Fatalf("could not read input file: %v", err)
	}

	r := bytes.NewReader(input)
	w := io.Discard

	for n := 0; n < b.N; n++ {
		r.Reset(input)
		_ = Solve(r, w)
	}
}

// go test -bench=. -cpu=1,2,4,8 -benchmem
// go test -bench='BenchmarkNoMemoryAlloc' -cpu=1,2,4,8 -benchmem

func BenchmarkMemoryMapAlloc(b *testing.B) {
	input, err := os.ReadFile("testdata/input.txt")
	if err != nil {
		b.Fatalf("could not read input file: %v", err)
	}

	r := bytes.NewReader(input)
	w := io.Discard

	for n := 0; n < b.N; n++ {
		r.Reset(input)
		_ = ASolve(r, w)
	}
}

// go test -bench='BenchmarkMemoryMapAlloc' -cpu=1,2,4,8 -benchmem

// 生成cpu profile 文件
// go test -bench='BenchmarkMemoryMapAlloc' -cpu=1,2,4,8 -cpuprofile=cpu.out
// go tool pprof -focus=Solve -call_tree -relative_percentages -png -output=cpu.png cpu.out
