package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// http://127.0.0.1:8899/check?path=D:\\code.txt
// https://www.cnblogs.com/vipsoft/p/17239713.html

func main() {
	log.Println("args:", os.Args)
	addr := "8899"
	if len(os.Args) > 1 {
		addr = os.Args[1]
	}
	log.Println("HTTP_PORT :", addr)
	server := http.NewServeMux()
	server.HandleFunc("GET /check", func(w http.ResponseWriter, r *http.Request) {
		var path = r.URL.Query().Get("path")
		log.Println("/check", path)
		exist, err := GetFileExist(path)
		w.Write([]byte(fmt.Sprintf("path:%v exist:%v err:%v", path, exist, err)))
		//w.WriteHeader(http.StatusOK)
	})
	http.ListenAndServe("0.0.0.0:"+addr, server)

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
