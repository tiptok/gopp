//+build wireinject

package main

import "github.com/google/wire"

func InitDb(dns string) (DB, error) {
	wire.Build(NewPostgresqlDB, NewDbConfig)
	return nil, nil
}

func InjectFooBar() FooBar {
	wire.Build(FooBarSet)
	return FooBar{}
}

//func InitDb(dns string)(DB,error){
//	wire.Build(DBSet)
//	return nil,nil
//}
