package resolvers

import (
	"github.com/graphql-go/graphql"
	"github.com/tiptok/gocomm/common"
	"github.com/tiptok/gopp/graphql/gographql/model"
	"github.com/tiptok/gopp/orm/pgorm/pg"
	"github.com/tiptok/gopp/orm/pgorm/pg/repository"
	"github.com/tiptok/gopp/orm/pgorm/pg/transaction"
	"github.com/tiptok/gopp/pkg/domain"
)

var UserResolvers = &userResolvers{}

type userResolvers struct {
}

func (resolvers *userResolvers) Roles(p graphql.ResolveParams) (interface{}, error) {
	if p.Source == nil {
		return nil, nil
	}
	if user, ok := p.Source.(*domain.Users); ok {
		return []model.Role{
			{Id: 1, RoleName: "运维" + user.Name},
		}, nil
	}
	return []model.Role{
		{Id: 1, RoleName: "运维"},
	}, nil
}

func (resolvers *userResolvers) Users(p graphql.ResolveParams) (interface{}, error) {
	transaction := transaction.NewGormTransactionContext(pg.DB)
	usersRepository, _ := repository.NewGormUsersRepository(transaction)
	total, users, err := usersRepository.Find(p.Args)
	return model.ListUser{
		Total: int(total),
		Users: users,
	}, err
}

func (resolvers *userResolvers) User(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(int)
	transaction := transaction.NewGormTransactionContext(pg.DB)
	usersRepository, _ := repository.NewGormUsersRepository(transaction)
	return usersRepository.FindOne(map[string]interface{}{"id": id})
}

func (resolvers *userResolvers) RemoveUser(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(int)
	transaction := transaction.NewGormTransactionContext(pg.DB)
	usersRepository, _ := repository.NewGormUsersRepository(transaction)
	transaction.StartTransaction()
	var err error
	var rsp *domain.Users
	defer func() {
		if err != nil {
			transaction.RollbackTransaction()
			return
		}
		transaction.CommitTransaction()
	}()
	rsp, err = usersRepository.Remove(&domain.Users{Id: int64(id)})
	return rsp, err
}

func (resolvers *userResolvers) UpdateUser(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(int)
	transaction := transaction.NewGormTransactionContext(pg.DB)
	usersRepository, _ := repository.NewGormUsersRepository(transaction)
	transaction.StartTransaction()
	var err error
	var rsp *domain.Users
	defer func() {
		if err != nil {
			transaction.RollbackTransaction()
			return
		}
		transaction.CommitTransaction()
	}()
	rsp, err = usersRepository.FindOne(map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	rsp.Update(p.Args)
	return usersRepository.Save(rsp)
}

func (resolvers *userResolvers) CreateUser(p graphql.ResolveParams) (interface{}, error) {
	transaction := transaction.NewGormTransactionContext(pg.DB)
	usersRepository, _ := repository.NewGormUsersRepository(transaction)
	transaction.StartTransaction()
	var err error
	var rsp *domain.Users
	defer func() {
		if err != nil {
			transaction.RollbackTransaction()
			return
		}
		transaction.CommitTransaction()
	}()
	var user = &domain.Users{}
	common.JsonUnmarshal(common.JsonAssertString(p.Args), user)
	rsp, err = usersRepository.Save(user)
	if err != nil {
		return rsp, err
	}
	return rsp, err
}
