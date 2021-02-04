package repository

import (
	"fmt"
	"github.com/tiptok/gocomm/common"
	"github.com/tiptok/gocomm/pkg/cache"
	. "github.com/tiptok/gocomm/pkg/orm/pgx"
	"github.com/tiptok/gopp/orm/go-pg/models"
	"github.com/tiptok/gopp/orm/go-pg/transaction"
	"github.com/tiptok/gopp/pkg/domain"
)

var (
	cacheUsersIdKey = func(id int64) string {
		//return fmt.Sprintf("%v:cache:Users:id:%v", constant.POSTGRESQL_DB_NAME, id)
		return ""
	}
)

type UsersRepository struct {
	transactionContext *transaction.TransactionContext
	*cache.CachedRepository
}

func (repository *UsersRepository) Save(dm *domain.Users) (*domain.Users, error) {
	var (
		err error
		m   = &models.Users{}
		tx  = repository.transactionContext.PgTx
	)
	if err = common.GobModelTransform(m, dm); err != nil {
		return nil, err
	}
	if dm.Identify() == nil {
		if err = tx.Insert(m); err != nil {
			return nil, err
		}
		dm.Id = m.Id
		return dm, nil
	}
	queryFunc := func() (interface{}, error) {
		return nil, tx.Update(m)
	}
	if _, err = repository.Query(queryFunc, cacheUsersIdKey(dm.Id)); err != nil {
		return nil, err
	}
	return dm, nil
}

func (repository *UsersRepository) Remove(User *domain.Users) (*domain.Users, error) {
	var (
		tx        = repository.transactionContext.PgTx
		UserModel = &models.Users{Id: User.Identify().(int64)}
	)
	queryFunc := func() (interface{}, error) {
		return tx.Model(UserModel).Where("id = ?", User.Id).Delete()
	}
	if _, err := repository.Query(queryFunc, cacheUsersIdKey(User.Id)); err != nil {
		return User, err
	}
	return User, nil
}

func (repository *UsersRepository) FindOne(queryOptions map[string]interface{}) (*domain.Users, error) {
	tx := repository.transactionContext.PgDd
	UserModel := new(models.Users)
	queryFunc := func() (interface{}, error) {
		query := NewQuery(tx.Model(UserModel), queryOptions)
		query.SetWhere("id = ?", "id")
		query.SetWhere("phone = ?", "phone")
		if err := query.First(); err != nil {
			return nil, fmt.Errorf("query row not found")
		}
		return UserModel, nil
	}
	var options []cache.QueryOption
	if _, ok := queryOptions["id"]; !ok {
		options = append(options, cache.WithNoCacheFlag())
	} else {
		UserModel.Id = queryOptions["id"].(int64)
	}
	if err := repository.QueryCache(cacheUsersIdKey(UserModel.Id), UserModel, queryFunc, options...); err != nil {
		return nil, err
	}

	if UserModel.Id == 0 {
		return nil, fmt.Errorf("query row not found")
	}
	return repository.transformPgModelToDomainModel(UserModel)
}

func (repository *UsersRepository) Find(queryOptions map[string]interface{}) (int64, []*domain.Users, error) {
	tx := repository.transactionContext.PgDd
	var UserModels []*models.Users
	Users := make([]*domain.Users, 0)
	query := NewQuery(tx.Model(&UserModels), queryOptions).
		SetOrder("create_time", "sortByCreateTime").
		SetOrder("update_time", "sortByUpdateTime").
		SetOrder("id", "sortById").
		SetLimit()

	if searchByText, ok := queryOptions["searchByText"]; ok && len(searchByText.(string)) > 0 {
		query.Where(fmt.Sprintf(`name like '%%%v%%'`, searchByText))
	}

	var err error
	if query.AffectRow, err = query.SelectAndCount(); err != nil {
		return 0, Users, err
	}
	for _, UserModel := range UserModels {
		if User, err := repository.transformPgModelToDomainModel(UserModel); err != nil {
			return 0, Users, err
		} else {
			Users = append(Users, User)
		}
	}
	return int64(query.AffectRow), Users, nil
}

func (repository *UsersRepository) transformPgModelToDomainModel(UserModel *models.Users) (*domain.Users, error) {
	m := &domain.Users{}
	err := common.GobModelTransform(m, UserModel)
	return m, err
}

func NewUserRepository(transactionContext *transaction.TransactionContext) (*UsersRepository, error) {
	if transactionContext == nil {
		return nil, fmt.Errorf("transactionContext参数不能为nil")
	}
	return &UsersRepository{transactionContext: transactionContext, CachedRepository: cache.NewDefaultCachedRepository()}, nil
}
