package main

import (
	"expvar"
	_ "expvar"
	"net/http"
)

func main() {
	expvar.NewMap()
	http.ListenAndServe(":8080", nil)
}
