// version_dev.go
// +build prod

package main

import (
	_ "embed"
)

//go:embed version.txt
var version string

//go:embed version.txt
var versionByte []byte

/*
测试
$ go run .
Version "dev"

$ go run -tags prod .
Version "0.0.1"
*/
