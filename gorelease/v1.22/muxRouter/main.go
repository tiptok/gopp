package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// 匹配方法
	mux.HandleFunc("POST /user/create", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "POST")
	})
	mux.HandleFunc("GET /user/get", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET")
	})
	// 通配符
	mux.HandleFunc("/items/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "id值为%s", id)
	})
	mux.HandleFunc("/items/{path...}", func(w http.ResponseWriter, r *http.Request) {
		path := r.PathValue("path")
		fmt.Fprintf(w, "path值为%s", path)
	})
	mux.HandleFunc("/items/{$}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "path /items/")
	})

	http.ListenAndServe(":8080", mux)
}
