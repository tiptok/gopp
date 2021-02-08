package main

import (
	"fmt"
	"github.com/google/wire"
)

var DBSet = wire.NewSet(NewDbConfig, NewMssqlDB, wire.Bind(new(DB), new(MssqlDB)))

type DB interface {
	Open() error
}

type MysqlDB struct {
	config DBConfig
}

func NewMysqlDB(config DBConfig) DB {
	return MysqlDB{
		config: config,
	}
}

func (db MysqlDB) Open() error {
	fmt.Println("mysql db open " + db.config)
	return nil
}

type PostgresqlDB struct {
	config DBConfig
}

func NewPostgresqlDB(config DBConfig) DB {
	return PostgresqlDB{
		config: config,
	}
}

func (db PostgresqlDB) Open() error {
	fmt.Println("Postgresql db open " + db.config)
	return nil
}

type MssqlDB struct {
	config DBConfig
}

func NewMssqlDB(config DBConfig) MssqlDB {
	return MssqlDB{
		config: config,
	}
}

func (db MssqlDB) Open() error {
	fmt.Println("mssql db open " + db.config)
	return nil
}
