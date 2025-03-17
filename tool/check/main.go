package main

import (
	"fmt"
	"net/http"
	"os"
)

// http://127.0.0.1:8899/check?path=D:\\code.txt
// https://www.cnblogs.com/vipsoft/p/17239713.html

func main() {
	//fmt.Println("args:", os.Args)
	//path := ""
	//if len(os.Args) > 1 {
	//	path = os.Args[1]
	//}
	//fmt.Println("check path:", path)
	server := http.NewServeMux()
	server.HandleFunc("GET /check", func(w http.ResponseWriter, r *http.Request) {
		var path = r.URL.Query().Get("path")
		exist, err := GetFileExist(path)
		w.Write([]byte(fmt.Sprintf("path:%v exist:%v err:%v", path, exist, err)))
		w.WriteHeader(http.StatusOK)
	})
	http.ListenAndServe(":8899", server)

}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.stat err:", err)
		bl := os.IsExist(err)
		fmt.Println("os.IsExist:", bl)
		if bl {
			return true
		}
		return false
	}
	return true
}

func GetFileExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
