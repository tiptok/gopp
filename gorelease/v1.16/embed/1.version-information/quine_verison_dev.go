//+build !prod

package main

import (
	_ "embed"
)

//go:embed version_dev.go
var versionSourceCode string
