//+build prod

package main

import (
	_ "embed"
)

//go:embed version_prd.go
var versionSourceCode string
