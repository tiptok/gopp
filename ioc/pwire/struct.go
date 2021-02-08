package main

import "github.com/google/wire"

type Foo int
type Bar int

func ProvideFoo() Foo {
	return Foo(1)
}

func ProvideBar() Bar {
	return Bar(2)
}

type FooBar struct {
	MyFoo Foo
	MyBar Bar
}

var FooBarSet = wire.NewSet(ProvideFoo, ProvideBar, wire.Struct(new(FooBar), "*")) // "MyFoo", "MyBar"
