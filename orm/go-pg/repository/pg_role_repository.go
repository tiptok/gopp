package repository

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/tiptok/gocomm/common"
	"github.com/tiptok/gocomm/pkg/cache"
	. "github.com/tiptok/gocomm/pkg/orm/pgx"
	"github.com/tiptok/gopp/orm/go-pg/models"
	"github.com/tiptok/gopp/orm/go-pg/transaction"
	"github.com/tiptok/gopp/pkg/constant"
	"github.com/tiptok/gopp/pkg/domain"
)

var (
	cacheRoleIdKey = func(v interface{}) string {
		return fmt.Sprintf("%v:cache:Role:id:%v", constant.POSTGRESQL_DB_NAME, v)
	}
)

type RoleRepository struct {
	*cache.CachedRepository
	transactionContext *transaction.TransactionContext
}

func (repository *RoleRepository) Save(dm *domain.Role) (*domain.Role, error) {
	var (
		err error
		m   = &models.Role{}
		tx  = repository.transactionContext.PgTx
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
		return tx.Exec(`update role set role_name=?,parent_id=?,update_time=now() where id = ?`, m.RoleName, m.ParentId, m.Id)
	}
	if _, err = repository.Query(queryFunc, cacheRoleIdKey(dm.Id)); err != nil {
		return nil, err
	}
	return dm, nil
}

func (repository *RoleRepository) Remove(Role *domain.Role) (*domain.Role, error) {
	var (
		tx        = repository.transactionContext.PgTx
		RoleModel = &models.Role{Id: Role.Identify().(int64)}
	)
	queryFunc := func() (interface{}, error) {
		return tx.Model(RoleModel).Where("id = ?", Role.Id).Delete()
	}
	if _, err := repository.Query(queryFunc, cacheRoleIdKey(Role.Id)); err != nil {
		return Role, err
	}
	return Role, nil
}

func (repository *RoleRepository) FindOne(queryOptions map[string]interface{}) (*domain.Role, error) {
	tx := repository.transactionContext.PgDd
	RoleModel := new(models.Role)
	queryFunc := func() (interface{}, error) {
		query := NewQuery(tx.Model(RoleModel), queryOptions)
		query.SetWhere("id = ?", "id")
		if err := query.First(); err != nil {
			return nil, fmt.Errorf("query row not found")
		}
		return RoleModel, nil
	}
	var options []cache.QueryOption
	var cacheModel = new(models.Role)
	if v, ok := queryOptions["id"]; ok {
		cacheModel.Id = v.(int64)
	}
	if err := repository.QueryCache(cacheModel.CacheKeyFunc, RoleModel, queryFunc, options...); err != nil {
		return nil, err
	}

	if RoleModel.Id == 0 {
		return nil, fmt.Errorf("query row not found")
	}
	return repository.transformPgModelToDomainModel(RoleModel)
}

func (repository *RoleRepository) Find(queryOptions map[string]interface{}) (int64, []*domain.Role, error) {
	tx := repository.transactionContext.PgDd
	var RoleModels []*models.Role
	Roles := make([]*domain.Role, 0)
	query := NewQuery(tx.Model(&RoleModels), queryOptions).
		SetWhere("parent_id=?", "parentId").
		SetOrder("create_time", "sortByCreateTime").
		SetOrder("update_time", "sortByUpdateTime").
		SetOrder("id", "orderById")

	var err error
	if inRoleIds, ok := queryOptions["inRoleIds"]; ok {
		query.Where("id in (?)", pg.In(inRoleIds))
	}
	if query.AffectRow, err = query.SelectAndCount(); err != nil {
		return 0, Roles, err
	}
	for _, RoleModel := range RoleModels {
		if Role, err := repository.transformPgModelToDomainModel(RoleModel); err != nil {
			return 0, Roles, err
		} else {
			Roles = append(Roles, Role)
		}
	}
	return int64(query.AffectRow), Roles, nil
}

func (repository *RoleRepository) transformPgModelToDomainModel(RoleModel *models.Role) (*domain.Role, error) {
	m := &domain.Role{}
	err := common.GobModelTransform(m, RoleModel)
	return m, err
}

func NewRoleRepository(transactionContext *transaction.TransactionContext) (*RoleRepository, error) {
	if transactionContext == nil {
		return nil, fmt.Errorf("transactionContext参数不能为nil")
	}
	return &RoleRepository{transactionContext: transactionContext, CachedRepository: cache.NewDefaultCachedRepository()}, nil
}
