package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// Reader
	s := strings.NewReader("ABCDEFG")
	br := bufio.NewReader(s) //1. NewReaderSize(rd, defaultBufSize)  4096

	// 2. Peek 返回缓存的一个切片，该切片引用缓存中前 n 字节数据
	//    该操作不会将数据读出，只是引用
	//    引用的数据在下一次读取操作之前是有效的
	b, _ := br.Peek(5)

	// 3. 修改切片的值 原来的值也变化了
	fmt.Println("before:", b, br.Buffered())
	b[0] = 66
	fmt.Println("after:", b, br.Buffered())

	// br.ReadByte
	// br.ReadBytes
	// br.ReadLine
	// br.ReadRune
	// br.ReadSlice
	// br.ReadString
	//br.Reset

	// Writer
	bf := bytes.NewBuffer(make([]byte, 0))
	br.WriteTo(bf)
	bw := bufio.NewWriter(bf) //bufio.NewReaderSize defaultBufSize 4096
	fmt.Println(bw.Available(), bw.Buffered())
	bw.WriteString(" hello 下午")

	bw.Flush()
	fmt.Println(bf.String())

	// Scanner
	sc := bufio.NewScanner(bytes.NewBufferString("aa bb cc\ndd"))
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
	if sc.Err() != nil {
		fmt.Println(sc.Err())
	}
}
