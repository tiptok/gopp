package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/tiptok/gocomm/common"
	"github.com/tiptok/gopp/orm/pgorm/pg"
	"github.com/tiptok/gopp/orm/pgorm/pg/repository"
	"github.com/tiptok/gopp/orm/pgorm/pg/transaction"
	"time"

	"github.com/Laisky/laisky-blog-graphql/types"
	"github.com/tiptok/gopp/graphql/pgqlgen/generated"
	"github.com/tiptok/gopp/graphql/pgqlgen/model"
	domain1 "github.com/tiptok/gopp/pkg/domain"
	"github.com/tiptok/gopp/pkg/protocol/user"
)

func (r *mutationResolver) CreateUsers(ctx context.Context, input *user.CreateUserRequest) (*domain1.Users, error) {
	transaction := transaction.NewGormTransactionContext(pg.DB)
	usersRepository, _ := repository.NewGormUsersRepository(transaction)
	transaction.StartTransaction()
	var err error
	var rsp *domain1.Users
	defer func() {
		if err != nil {
			transaction.RollbackTransaction()
			return
		}
		transaction.CommitTransaction()
	}()
	rsp, err = usersRepository.Save(&domain1.Users{
		Name:       input.Name,
		Passwd:     input.Passwd,
		Phone:      input.Phone,
		Roles:      input.Roles,
		AdminType:  input.AdminType,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	if err != nil {
		return rsp, err
	}
	return rsp, err
}

func (r *mutationResolver) RemoveUsers(ctx context.Context, input *user.DeleteUserRequest) (*domain1.Users, error) {
	transaction := transaction.NewGormTransactionContext(pg.DB)
	usersRepository, _ := repository.NewGormUsersRepository(transaction)
	transaction.StartTransaction()
	var err error
	var rsp *domain1.Users
	defer func() {
		if err != nil {
			transaction.RollbackTransaction()
			return
		}
		transaction.CommitTransaction()
	}()
	rsp, err = usersRepository.Remove(&domain1.Users{Id: input.Id})
	return rsp, err
}

func (r *mutationResolver) UpdateUsers(ctx context.Context, input *user.UpdateUserRequest) (*domain1.Users, error) {
	transaction := transaction.NewGormTransactionContext(pg.DB)
	usersRepository, _ := repository.NewGormUsersRepository(transaction)
	transaction.StartTransaction()
	var err error
	var rsp *domain1.Users
	defer func() {
		if err != nil {
			transaction.RollbackTransaction()
			return
		}
		transaction.CommitTransaction()
	}()
	rsp, err = usersRepository.Save(&domain1.Users{Id: input.Id})
	return rsp, err
}

func (r *queryResolver) User(ctx context.Context, input *user.GetUserRequest) (*domain1.Users, error) {
	transaction := transaction.NewGormTransactionContext(pg.DB)
	usersRepository, _ := repository.NewGormUsersRepository(transaction)
	return usersRepository.FindOne(map[string]interface{}{"id": input.Id})
}

func (r *queryResolver) Users(ctx context.Context, input *user.ListUserRequest) (*model.ListUser, error) {
	transaction := transaction.NewGormTransactionContext(pg.DB)
	usersRepository, _ := repository.NewGormUsersRepository(transaction)
	total, users, err := usersRepository.Find(common.ObjectToMap(input))
	return &model.ListUser{
		Total: int(total),
		Users: users,
	}, err
}

func (r *usersResolver) Roles(ctx context.Context, obj *domain1.Users) ([]*model.Role, error) {
	return []*model.Role{&model.Role{}}, nil
}

func (r *usersResolver) CreateTime(ctx context.Context, obj *domain1.Users) (*types.Datetime, error) {
	return types.NewDatetimeFromTime(obj.CreateTime), nil
}

func (r *usersResolver) UpdateTime(ctx context.Context, obj *domain1.Users) (*types.Datetime, error) {
	return types.NewDatetimeFromTime(obj.UpdateTime), nil
}

// Users returns generated.UsersResolver implementation.
func (r *Resolver) Users() generated.UsersResolver { return &usersResolver{r} }

type usersResolver struct{ *Resolver }
