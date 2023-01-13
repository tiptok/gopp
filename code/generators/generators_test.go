package main

import "testing"

func TestStringer(t *testing.T) {
	status := Verified
	t.Log(status.String())
}
