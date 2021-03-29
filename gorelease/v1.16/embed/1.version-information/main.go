package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var (
	Version string = strings.TrimSpace(version)
)

func main() {
	fmt.Printf("Version %q\n", Version)
	fmt.Printf("Version Byte %s\n", string(versionByte))
	fmt.Print("\nsource code:\n", versionSourceCode)
}
