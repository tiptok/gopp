package transaction

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"sync"
)

type TransactionContext struct {
	//启用事务标识
	beginTransFlag bool
	rawDb          *pg.DB
	session        orm.DB
	lock           sync.Mutex
}

func (transactionContext *TransactionContext) Begin() error {
	transactionContext.lock.Lock()
	defer transactionContext.lock.Unlock()
	transactionContext.beginTransFlag = true
	tx, err := transactionContext.rawDb.Begin()
	if err != nil {
		return err
	}
	transactionContext.session = tx
	return nil
}

func (transactionContext *TransactionContext) Commit() error {
	transactionContext.lock.Lock()
	defer transactionContext.lock.Unlock()
	if !transactionContext.beginTransFlag {
		return nil
	}
	if v, ok := transactionContext.session.(*pg.Tx); ok {
		err := v.Commit()
		return err
	}
	return nil
}

func (transactionContext *TransactionContext) Rollback() error {
	transactionContext.lock.Lock()
	defer transactionContext.lock.Unlock()
	if !transactionContext.beginTransFlag {
		return nil
	}
	if v, ok := transactionContext.session.(*pg.Tx); ok {
		err := v.Rollback()
		return err
	}
	return nil
}

func (transactionContext *TransactionContext) DB() orm.DB {
	return transactionContext.session
}

func NewTransactionContext(db *pg.DB) *TransactionContext {
	return &TransactionContext{
		rawDb:   db,
		session: db,
	}
}
