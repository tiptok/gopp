package main

import (
	"log"
	"net/http"
	"net/http/cgi"
	"testing"
)

//cgi
//http请求 通过cgi调用指定cmd 返回结果
func Test_CGI_Server(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler := new(cgi.Handler)
		handler.Path = defaultEnv("GOROOT", "") + "/bin/go.exe"
		script := `F:\go\src\github.com\tiptok\OFAppTest\src\golib\netT\httpT` + r.URL.Path
		log.Println(handler.Path)
		handler.Dir = ""
		args := []string{"run", script}
		handler.Args = append(handler.Args, args...)
		handler.Env = append(handler.Env, "GOPATH="+defaultEnv("GOPATH", ""))
		handler.Env = append(handler.Env, "GOROOT="+defaultEnv("GOROOT", ""))
		log.Println(handler.Args)

		handler.ServeHTTP(w, r)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func defaultEnv(k, d string) string {
	return ""
}

//http://127.0.0.1:8080/cgi-script.go
