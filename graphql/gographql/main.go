package main

import (
	"github.com/tiptok/gopp/graphql/gographql/schemal"
	"net/http"
)

const defaultPort = "8080"

func main() {
	http.Handle("/query", schemal.InitHandler())
	http.ListenAndServe(":"+defaultPort, nil)
}
