package byteT

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNewBuffer(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	n, err := buf.Write([]byte("hello 2018"))
	if err == nil {
		t.Log("Write buffer size:", n)
		fmt.Println("Write buffer size:", n)
	}
	br, _ := buf.ReadByte()
	fmt.Println("Read byte:", br)
	PrintBufInfo(buf)

	buf.WriteRune('新')
	buf.WriteRune('年')
	PrintBufInfo(buf)

	buf.WriteString("新开始")
	PrintBufInfo(buf)

	temp := buf.Bytes()
	//buf.Reset()   == buf.Truncate(0)
	buf.Truncate(0)
	PrintBufInfo(buf)

	buf.Write(temp)
	PrintBufInfo(buf)

	// bufRNew := bytes.NewBuffer(nil)
	// bRFrom, _ := bufRNew.ReadFrom(buf)
	// fmt.Printf("ReadFrom Buffer Size:%d\n", bRFrom)
	//读完 buffer就没有数据

	bWToNew := bytes.NewBufferString("H")
	bWTo, _ := buf.WriteTo(bWToNew)
	fmt.Printf("Write To Buffer Size:%d\n", bWTo)
	PrintBufInfo(bWToNew)
}

func PrintBufInfo(buf *bytes.Buffer) {
	fmt.Println("Current Buffer Len:", buf.Len(), "Left Buffer:", string(buf.Bytes()))
}
