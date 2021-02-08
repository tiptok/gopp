package main

import "github.com/tiptok/gocomm/pkg/log"

func main() {
	db, err := InitDb("127.0.0.1:3306")
	if err != nil {
		log.Error(err)
	}
	db.Open()
}
