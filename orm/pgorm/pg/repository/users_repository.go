package repository

import (
	"fmt"
	"github.com/tiptok/gocomm/common"
	"github.com/tiptok/gopp/orm/pgorm/pg/models"
	"github.com/tiptok/gopp/orm/pgorm/pg/transaction"
	"github.com/tiptok/gopp/pkg/domain"
)

type GormUsersRepository struct {
	transactionContext *transaction.TransactionContext
}

func (repository *GormUsersRepository) Save(dm *domain.Users) (*domain.Users, error) {
	var (
		err error
		m   = &models.Users{}
		tx  = repository.transactionContext.PgTx
	)
	if err = common.GobModelTransform(m, dm); err != nil {
		return nil, err
	}
	m.Roles = dm.Roles
	if tx = tx.Save(m); tx.Error != nil {
		return nil, tx.Error
	}
	dm.Id = m.Id
	return dm, nil
}

func (repository *GormUsersRepository) Remove(Users *domain.Users) (*domain.Users, error) {
	var (
		tx         = repository.transactionContext.PgTx
		UsersModel = &models.Users{Id: Users.Identify().(int64)}
	)
	if tempTx := tx.Where("id = ?", Users.Id).Delete(UsersModel); tempTx.Error != nil {
		return Users, tempTx.Error
	}
	return Users, nil
}

func (repository GormUsersRepository) FindOne(options map[string]interface{}) (*domain.Users, error) {
	tx := repository.transactionContext.PgDb
	m := &models.Users{}
	if v, ok := options["id"]; ok {
		tx = tx.Where("id = ?", v)
	}
	tempTx := tx.Find(m)
	if tempTx.Error != nil {
		return nil, tempTx.Error
	}
	return repository.transformPgModelToDomainModel(m)
}

func (repository *GormUsersRepository) Find(queryOptions map[string]interface{}) (int64, []*domain.Users, error) {
	var (
		tx          = repository.transactionContext.PgDb
		err         error
		UsersModels []*models.Users
		users       = make([]*domain.Users, 0)
	)
	tempTx := tx.Find(&UsersModels)
	if tempTx.Error != nil {
		return 0, users, err
	}
	for _, UsersModel := range UsersModels {
		if Users, err := repository.transformPgModelToDomainModel(UsersModel); err != nil {
			return 0, users, err
		} else {
			users = append(users, Users)
		}
	}
	return tempTx.RowsAffected, users, nil
}

func (repository *GormUsersRepository) transformPgModelToDomainModel(UsersModel *models.Users) (*domain.Users, error) {
	m := &domain.Users{}
	err := common.GobModelTransform(m, UsersModel)
	m.Roles = UsersModel.Roles
	return m, err
}

func NewGormUsersRepository(transactionContext *transaction.TransactionContext) (*GormUsersRepository, error) {
	if transactionContext == nil {
		return nil, fmt.Errorf("transactionContext参数不能为nil")
	}
	return &GormUsersRepository{transactionContext: transactionContext}, nil
}
