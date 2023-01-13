package main

//go:generate mockery  --name=SendFunc

type SendFunc func(data string) (int, error)

//go:generate mockery  --name=Stringer
type Stringer interface {
	String() string
}

// go install github.com/vektra/mockery/v2@latest
// go generate ./...
