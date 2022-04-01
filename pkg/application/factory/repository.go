package factory

import (
	"github.com/tiptok/gopp/orm/go-pg/repository"
	"github.com/tiptok/gopp/orm/go-pg/transaction"
	"github.com/tiptok/gopp/pkg/domain"
)

func CreateRoleRepository(transactionContext *transaction.TransactionContext) (domain.RoleRepository, error) {
	return repository.NewRoleRepository(transactionContext)
}

func CreateUserRepository(transactionContext *transaction.TransactionContext) (domain.UsersRepository, error) {
	return repository.NewUserRepository(transactionContext)
}
