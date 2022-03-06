package ioT

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	// destFile, err := os.Create("temp.txt")
	// if err != nil {
	// 	fmt.Println("Open File Error:", err.Error())
	// }
	// destFile.WriteString("2018-02-28 os.Create")

	opFile, err := os.OpenFile("temp.txt", os.O_RDWR, 666)
	/*os.Open -> readonly */
	if err != nil {
		fmt.Println("Open File Error:", err.Error())
	}
	opFile.WriteString("\n tik tik tik")
	opFile.WriteString("\n tok tok tok")
	t.Log("End")
}

func Test_CopyN(t *testing.T) {
	var n int64
	file, err := os.OpenFile("temp.txt", os.O_RDWR, 666)
	if err != nil {
		t.Fatal(err)
	}
	var buf *bytes.Buffer = bytes.NewBuffer(nil)
	if n, err = io.CopyN(buf, file, 10); err != nil {
		t.Fatal(err)
	}
	t.Log(buf.String(), " length:", n)

	if tmpFile, err := ioutil.TempFile("", "copyn_"); err != nil {

	} else {
		defer tmpFile.Close()
		tmpFile.Write(buf.Bytes())
		t.Log(tmpFile.Name())
	}
}

func Test_SectionReader(t *testing.T) {
	file, err := os.OpenFile("temp.txt", os.O_RDWR, 666)
	if err != nil {
		t.Fatal(err)
	}
	r := io.NewSectionReader(file, 10, 5)
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, r)
	t.Log(buf.String())
}
