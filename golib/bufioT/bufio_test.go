package main

import (
	"bufio"
	"bytes"
	"log"
	"testing"
)

func Test_Scanner(t *testing.T) {
	s := bufio.NewScanner(bytes.NewBufferString("aa bb cc"))
	s.Split(bufio.ScanWords) //bufio.ScanLines
	for s.Scan() {
		log.Println(s.Text())
	}
	if s.Err() != nil {
		t.Fatal(s.Err())
	}
	//s.Buffer(nil,bufio.MaxScanTokenSize)
}
