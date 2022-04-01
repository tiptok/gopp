package factory

import (
	go_pg "github.com/tiptok/gopp/orm/go-pg"
	"github.com/tiptok/gopp/orm/go-pg/transaction"
)

func CreateTransactionContext() *transaction.TransactionContext {
	return transaction.NewTransactionContext(go_pg.DB)
}
