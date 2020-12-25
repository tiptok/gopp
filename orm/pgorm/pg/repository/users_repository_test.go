package repository

import (
	"github.com/tiptok/gocomm/common"
	"github.com/tiptok/gopp/orm/pgorm/pg"
	_ "github.com/tiptok/gopp/orm/pgorm/pg"
	"github.com/tiptok/gopp/orm/pgorm/pg/transaction"
	"github.com/tiptok/gopp/pkg/domain"
	"testing"
	"time"
)

var users = &domain.Users{
	Id:         1,
	Name:       "tiptok",
	Phone:      "18860183060",
	Passwd:     "123456",
	Roles:      []int64{6, 8},
	Status:     1,
	AdminType:  1,
	CreateTime: time.Now(),
	UpdateTime: time.Now(),
}

func TestGormUsersRepository_Save(t *testing.T) {
	tx := transaction.NewGormTransactionContext(pg.DB)
	var err error
	usersRepository, _ := NewGormUsersRepository(tx)
	defer func() {
		tx.RollbackTransaction()
	}()
	err = tx.StartTransaction()
	if err != nil {
		return
	}
	u := *users
	u.Id = 0
	var user = &u
	// Save
	if user, err = usersRepository.Save(user); err != nil {
		return
	}
	if err = tx.CommitTransaction(); err != nil {
		return
	}
	//if _,err = usersRepository.Remove(user);err!=nil{
	//	return
	//}
}

func TestGormUsersRepository_FindOne(t *testing.T) {
	tx := transaction.NewGormTransactionContext(pg.DB)
	var err error
	usersRepository, _ := NewGormUsersRepository(tx)
	defer func() {
		tx.RollbackTransaction()
	}()
	err = tx.StartTransaction()
	if err != nil {
		return
	}
	var user *domain.Users
	if user, err = usersRepository.FindOne(map[string]interface{}{"id": 1}); err != nil {
		return
	}
	if err = tx.CommitTransaction(); err != nil {
		return
	}
	if user.Id != 1 {
		t.Fatal("error id :", user.Id)
	}
	//t.Log(common.JsonAssertString(user))
}

func TestGormUsersRepository_Remove(t *testing.T) {
	tx := transaction.NewGormTransactionContext(pg.DB)
	var err error
	usersRepository, _ := NewGormUsersRepository(tx)
	defer func() {
		tx.RollbackTransaction()
	}()
	err = tx.StartTransaction()
	if err != nil {
		return
	}
	var user *domain.Users
	if user, err = usersRepository.Remove(users); err != nil {
		return
	}
	if err = tx.CommitTransaction(); err != nil {
		return
	}
	if user.Id != 1 {
		t.Fatal("error id :", user.Id)
	}
}

func TestGormUsersRepository_Find(t *testing.T) {
	tx := transaction.NewGormTransactionContext(pg.DB)
	var err error
	usersRepository, _ := NewGormUsersRepository(tx)
	defer func() {
		tx.RollbackTransaction()
	}()
	err = tx.StartTransaction()
	if err != nil {
		return
	}
	if _, users, err := usersRepository.Find(nil); err != nil {
		return
	} else {
		t.Log(common.JsonAssertString(users))
	}
	if err = tx.CommitTransaction(); err != nil {
		return
	}
}
