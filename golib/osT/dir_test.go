package osT

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestReadDir(t *testing.T) {
	file, err := os.Open("D:")
	defer file.Close()
	if err != nil {
		t.Fatal(err)
	}
	fileInfos, err := file.Readdir(0)
	if err != nil {
		t.Fatal(err)
	}
	for i, fileInfo := range fileInfos {
		log.Println(fmt.Sprintf("Dir Index %v %v", i, fileInfo))
	}
	fileInfoName, err := file.Readdirnames(0)
	if err != nil {
		t.Fatal(err)
	}
	for i, fileInfo := range fileInfoName {
		log.Println(fmt.Sprintf("Dir Index %v %v", i, fileInfo))
	}
}

func TestEnv(t *testing.T) {
	buf := make([]byte, 0, 20)
	msg := "append data"
	buf = append(buf, msg...)
	fmt.Println(string(buf))

	mapping := func(s string) string {
		m := map[string]string{"xx": "sssssssssssss",
			"yy": "ttttttttttttttt"}
		return m[s]
	}
	datas := "hello $xx blog address $yy"
	expandStr := os.Expand(datas, mapping)
	fmt.Println(expandStr)
}
