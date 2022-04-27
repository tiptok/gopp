package repository

import (
	"fmt"
	"github.com/tiptok/gocomm/common"
	"github.com/tiptok/gocomm/pkg/cache"
	. "github.com/tiptok/gocomm/pkg/orm/pgx"
	"github.com/tiptok/gopp/orm/go-pg/models"
	"github.com/tiptok/gopp/orm/go-pg/transaction"
	"github.com/tiptok/gopp/pkg/constant"
	"github.com/tiptok/gopp/pkg/domain"
)

var (
	cacheUsersIdKey = func(v interface{}) string {
		return fmt.Sprintf("%v:cache:Users:id:%v", constant.POSTGRESQL_DB_NAME, v)
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
		tx  = repository.transactionContext.DB()
	)
	if err = common.GobModelTransform(m, dm); err != nil {
		return nil, err
	}
	if dm.Identify() == nil {
		if _, err = tx.Model(m).Insert(m); err != nil {
			return nil, err
		}
		dm.Id = m.Id
		return dm, nil
	}
	queryFunc := func() (interface{}, error) {
		return tx.Model(m).WherePK().Update(m)
	}
	if _, err = repository.Query(queryFunc, m.CacheKeyFunc()); err != nil {
		return nil, err
	}
	return dm, nil
}

func (repository *UsersRepository) Remove(User *domain.Users) (*domain.Users, error) {
	var (
		tx        = repository.transactionContext.DB()
		UserModel = &models.Users{Id: User.Identify().(int64)}
	)
	queryFunc := func() (interface{}, error) {
		return tx.Model(UserModel).Where("id = ?", User.Id).Delete()
	}
	if _, err := repository.Query(queryFunc, UserModel.CacheKeyFunc()); err != nil {
		return User, err
	}
	return User, nil
}

func (repository *UsersRepository) FindOne(queryOptions map[string]interface{}) (*domain.Users, error) {
	tx := repository.transactionContext.DB()
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
	cacheModel := new(models.Users)
	if _, ok := queryOptions["id"]; ok {
		cacheModel.Id = queryOptions["id"].(int64)
	}
	if err := repository.QueryCache(cacheModel.CacheKeyFunc, UserModel, queryFunc); err != nil {
		return nil, err
	}

	if UserModel.Id == 0 {
		return nil, fmt.Errorf("query row not found")
	}
	return repository.transformPgModelToDomainModel(UserModel)
}

func (repository *UsersRepository) FindOneByPhone(phone string) (*domain.Users, error) {
	tx := repository.transactionContext.DB()
	UserModel := new(models.Users)
	queryFunc := func() (interface{}, error) {
		query := NewQuery(tx.Model(UserModel), nil)
		query.Where("phone = ?", phone)
		if err := query.First(); err != nil {
			return nil, fmt.Errorf("query row not found")
		}
		return UserModel, nil
	}
	cacheModel := new(models.Users)
	cacheModel.Phone = phone
	if err := repository.QueryUniqueIndexCache(cacheModel.CachePrimaryKeyFunc, UserModel, func(o interface{}) string {
		if v, ok := o.(*models.Users); ok {
			return v.CacheKeyFunc()
		}
		return ""
	}, queryFunc); err != nil {
		return nil, err
	}

	if UserModel.Id == 0 {
		return nil, fmt.Errorf("query row not found")
	}
	return repository.transformPgModelToDomainModel(UserModel)
}

func (repository *UsersRepository) FindOneByPhoneNoCache(phone string) (*domain.Users, error) {
	tx := repository.transactionContext.DB()
	UserModel := new(models.Users)
	query := NewQuery(tx.Model(UserModel), nil)
	query.Where("phone = ?", phone)
	if err := query.First(); err != nil {
		return nil, fmt.Errorf("query row not found")
	}
	return repository.transformPgModelToDomainModel(UserModel)
}

func (repository *UsersRepository) Find(queryOptions map[string]interface{}) (int64, []*domain.Users, error) {
	tx := repository.transactionContext.DB()
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
