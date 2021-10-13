package ioT

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

/*
	NopCloser
*/
func TestNopCloser(t *testing.T) {
	s := strings.NewReader("hello 2018 to go")
	r := ioutil.NopCloser(s)
	sli := make([]byte, s.Len())
	nR, err := s.Read(sli)
	if err == nil {
		fmt.Println("Read bytes Num:", nR)
		fmt.Printf("Reader.Read String:%v\n", string(sli[:]))
	}
	fmt.Printf("Left String Len:%v\n", s.Len())
	t.Log("End")

	buf := bytes.NewBuffer(sli)
	b, err := ioutil.ReadAll(buf)
	fmt.Println("ioutil.ReadAll:", string(b))

	if err != nil && err.Error() == "EOF" {
		fmt.Println("Read To End.")
	}
	defer r.Close()
}

/*
	ioutil.ReadDir
	ioutil.ReadFile
*/

func TestReadDir(t *testing.T) {
	fmt.Println("Please Input Direct Root:")
	var root string = ""
	//fmt.Scanln(&root)
	if root == "" {
		/*root 为空 默认D://*/
		root = "D:/Go/src/github.com/tiptok/GoNas" //
	}
	ReadDir(root)
	ReadFile(root, "gonas.go")
	t.Log("End")
}

func PrintDirInfo(dirList []os.FileInfo) {
	var sFlag = "+"

	for i, v := range dirList {
		if v.IsDir() {
			sFlag = "+"
		} else {
			sFlag = " "
		}
		fmt.Printf("%d %s %v %v\n", i, sFlag, v.Name(), v.ModTime())
	}
}

func ReadDir(sPath string) {
	dirList, err := ioutil.ReadDir(sPath)
	if err == nil {
		PrintDirInfo(dirList)
	} else {
		fmt.Println("Read dir error.", err.Error())
	}
}

func ReadFile(sPath, sFileName string) {
	sFilePath := sPath + "/" + sFileName
	file, err := ioutil.ReadFile(sFilePath)
	if err != nil {
		fmt.Println("read file error:", err.Error())
	}
	fmt.Printf("\nReadFile:%s \n", sFilePath)
	fmt.Println(string(file))
}
