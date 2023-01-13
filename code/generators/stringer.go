package main

//go:generate stringer -type=UserStatus
type UserStatus int

const (
	Unverified UserStatus = iota
	Verified
	Suspended
	Unknown
)

// 安装stringer
// go install golang.org/x/tools/cmd/stringer
