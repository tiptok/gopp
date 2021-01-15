package transaction

import "github.com/go-pg/pg/v10"

type TransactionContext struct {
	CloseTransactionFlag bool //事务关闭标识
	PgDd                 *pg.DB
	PgTx                 *pg.Tx
}

func (transactionContext *TransactionContext) StartTransaction() error {
	if transactionContext.CloseTransactionFlag {
		return nil
	}
	tx, err := transactionContext.PgDd.Begin()
	if err != nil {
		return err
	}
	transactionContext.PgTx = tx
	return nil
}

func (transactionContext *TransactionContext) CommitTransaction() error {
	if transactionContext.CloseTransactionFlag {
		return nil
	}
	err := transactionContext.PgTx.Commit()
	return err
}

func (transactionContext *TransactionContext) RollbackTransaction() error {
	if transactionContext.CloseTransactionFlag {
		return nil
	}
	err := transactionContext.PgTx.Rollback()
	return err
}

// SetTransactionClose
// 在不需要事务的地方可以执行该方法关闭事务处理
// 例如:对象的查询
func (transactionContext *TransactionContext) SetTransactionClose() error {
	transactionContext.CloseTransactionFlag = true
	return nil
}

func NewPGTransactionContext(pgDd *pg.DB) *TransactionContext {
	return &TransactionContext{
		PgDd: pgDd,
	}
}
