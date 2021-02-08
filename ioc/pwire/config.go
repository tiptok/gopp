package main

import "github.com/google/wire"

type DBConfig string

func NewDbConfig(dns string) DBConfig {
	return DBConfig(dns)
}

var set = wire.NewSet()
