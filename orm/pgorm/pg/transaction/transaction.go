package transaction

import (
	"github.com/tiptok/gopp/orm/pgorm/pg"
	"gorm.io/gorm"
)

type TransactionContext struct {
	PgTx *gorm.DB
	PgDb *gorm.DB
}

func (transactionContext *TransactionContext) StartTransaction() error {
	tx := pg.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	transactionContext.PgTx = tx
	return nil
}

func (transactionContext *TransactionContext) CommitTransaction() error {
	tx := transactionContext.PgTx.Commit()
	return tx.Error
}

func (transactionContext *TransactionContext) RollbackTransaction() error {
	if transactionContext.PgTx.Error == nil {
		return nil
	}
	tx := transactionContext.PgTx.Rollback()
	return tx.Error
}

func NewGormTransactionContext(pgDd *gorm.DB) *TransactionContext {
	return &TransactionContext{
		PgDb: pgDd,
	}
}
