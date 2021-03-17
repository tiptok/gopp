package go_pg

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/tiptok/gopp/orm/go-pg/models"
	"github.com/tiptok/gopp/pkg/constant"
	"log"
)

var DB *pg.DB

func init() {
	DB = pg.Connect(&pg.Options{
		User:     constant.POSTGRESQL_USER,
		Password: constant.POSTGRESQL_PASSWORD,
		Database: constant.POSTGRESQL_DB_NAME,
		Addr:     fmt.Sprintf("%s:%s", constant.POSTGRESQL_HOST, constant.POSTGRESQL_PORT),
	})
	if !constant.DISABLE_SQL_GENERATE_PRINT {
		DB.AddQueryHook(SqlGeneratePrintHook{})
	}
	//orm.RegisterTable((*models.OrderGood)(nil))
	if !constant.DISABLE_CREATE_TABLE {
		for _, model := range []interface{}{
			(*models.Users)(nil),
			(*models.SysMessageConsume)(nil),
			(*models.SysMessageProduce)(nil),
		} {
			err := DB.Model(model).CreateTable(&orm.CreateTableOptions{
				Temp:          false,
				IfNotExists:   true,
				FKConstraints: true,
			})
			if err != nil {
				panic(err)
			}
		}
	}
}

type SqlGeneratePrintHook struct{}

func (hook SqlGeneratePrintHook) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (hook SqlGeneratePrintHook) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	data, err := q.FormattedQuery()
	//if len(string(data)) > 8 { //BEGIN COMMIT
	log.Println(string(data))
	//}
	return err
}
