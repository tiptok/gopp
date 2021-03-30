package user

import (
	"crypto/sha1"
	"fmt"
	"github.com/tiptok/gocomm/common"
	"github.com/tiptok/gocomm/gs"
	"github.com/tiptok/gocomm/pkg/log"
	go_pg "github.com/tiptok/gopp/orm/go-pg"
	"github.com/tiptok/gopp/orm/go-pg/repository"
	"github.com/tiptok/gopp/orm/go-pg/transaction"
	"github.com/tiptok/gopp/pkg/application/factory"
	"github.com/tiptok/gopp/pkg/domain"
	protocol "github.com/tiptok/gopp/pkg/protocol"
	protocolx "github.com/tiptok/gopp/pkg/protocol/user"
	"strings"
	"time"
)

type UserService struct {
}

func (svr *UserService) CreateUser(header *protocol.RequestHeader, request *protocolx.CreateUserRequest) (rsp interface{}, err error) {
	var (
		transactionContext = transaction.NewPGTransactionContext(go_pg.DB)
	)
	rsp = &protocolx.CreateUserResponse{}
	if err = request.ValidateCommand(); err != nil {
		//err = application.ThrowError(application.ARG_ERROR, err.Error())
		return
	}
	if err = transactionContext.StartTransaction(); err != nil {
		//err = application.ThrowError(application.TRANSACTION_ERROR, err.Error())
		return
	}
	defer func() {
		transactionContext.RollbackTransaction()
	}()
	newUser := &domain.Users{
		Name:       request.Name,
		Phone:      request.Phone,
		Roles:      request.Roles,
		Status:     1,
		AdminType:  request.AdminType,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	var UserRepository, _ = repository.NewUserRepository(transactionContext) //factory.CreateUserRepository(transactionContext)
	if request.Phone != "" {
		if _, err = UserRepository.FindOne(map[string]interface{}{"phone": request.Phone}); err == nil {
			err = protocol.NewCustomMessage(1, "手机号已存在")
			return
		}
	}
	if len(newUser.Passwd) == 0 {
		newUser.Passwd = fmt.Sprintf("%x", sha1.Sum([]byte("123456")))
	}
	if m, e := UserRepository.Save(newUser); e != nil {
		//err = application.ThrowError(application.INTERNAL_SERVER_ERROR, e.Error())
		return
	} else {
		rsp = m
	}
	if err = transactionContext.CommitTransaction(); err != nil {
		//err = application.ThrowError(application.TRANSACTION_ERROR, err.Error())
		return
	}
	return
}

func (svr *UserService) UpdateUser(header *protocol.RequestHeader, request *protocolx.UpdateUserRequest) (rsp interface{}, err error) {
	var (
		transactionContext = transaction.NewPGTransactionContext(go_pg.DB)
	)
	rsp = &protocolx.UpdateUserResponse{}
	if err = request.ValidateCommand(); err != nil {
		//err = application.ThrowError(application.ARG_ERROR, err.Error())
		return
	}
	if err = transactionContext.StartTransaction(); err != nil {
		log.Error(err)
		return nil, err
	}
	defer func() {
		transactionContext.RollbackTransaction()
	}()

	var UserRepository, _ = repository.NewUserRepository(transactionContext) //factory.CreateUserRepository(transactionContext)
	var user *domain.Users
	if user, err = UserRepository.FindOne(map[string]interface{}{"id": request.Id}); err != nil {
		err = protocol.NewCustomMessage(1, "用户不存在")
		return
	}
	if request.Phone != "" && request.Phone != user.Phone {
		if _, err = UserRepository.FindOne(map[string]interface{}{"phone": request.Phone}); err == nil {
			err = protocol.NewCustomMessage(1, "手机号已存在")
			return
		}
	}
	//  common.ObjectToMap(request)
	if err = user.Update(common.LoadCustomFieldToMap(request, header.BodyKeys...)); err != nil {
		//err = application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
		return
	}
	if user, err = UserRepository.Save(user); err != nil {
		//err = application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
		return
	}
	if err = transactionContext.CommitTransaction(); err != nil {
		//err = application.ThrowError(application.TRANSACTION_ERROR, err.Error())
		return
	}
	return
}

func (svr *UserService) GetUser(header *protocol.RequestHeader, request *protocolx.GetUserRequest) (rsp interface{}, err error) {
	var (
		transactionContext = transaction.NewPGTransactionContext(go_pg.DB)
		RoleRepository, _  = factory.CreateRoleRepository(transactionContext)
	)
	rsp = &protocolx.GetUserResponse{}
	if err = request.ValidateCommand(); err != nil {
		//err = application.ThrowError(application.ARG_ERROR, err.Error())
		return
	}
	if err = transactionContext.StartTransaction(); err != nil {
		//err = application.ThrowError(application.TRANSACTION_ERROR, err.Error())
		return
	}
	defer func() {
		transactionContext.RollbackTransaction()
	}()

	var UserRepository, _ = repository.NewUserRepository(transactionContext) // factory.CreateUserRepository(transactionContext)
	var user *domain.Users
	if user, err = UserRepository.FindOne(common.ObjectToMap(request)); err != nil {
		//err = application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
		return
	}
	retMap := map[string]interface{}{"id": user.Id, "name": user.Name, "phone": user.Phone, "adminType": user.AdminType, "status": user.Status}
	var roles []*domain.Role
	for _, v := range user.Roles {
		if role, e := RoleRepository.FindOne(map[string]interface{}{"id": v}); e == nil {
			roles = append(roles, role)
		}
	}
	retMap["roles"] = common.LoadCustomField(roles, "Id", "RoleName")
	rsp = (gs.MapData)(map[string]interface{}{"user": retMap})
	if err = transactionContext.CommitTransaction(); err != nil {
		//err = application.ThrowError(application.TRANSACTION_ERROR, err.Error())
		return
	}
	return
}

func (svr *UserService) DeleteUser(header *protocol.RequestHeader, request *protocolx.DeleteUserRequest) (rsp interface{}, err error) {
	var (
		transactionContext = transaction.NewPGTransactionContext(go_pg.DB)
	)
	rsp = &protocolx.DeleteUserResponse{}
	if err = request.ValidateCommand(); err != nil {
		//err = application.ThrowError(application.ARG_ERROR, err.Error())
		return
	}
	if err = transactionContext.StartTransaction(); err != nil {
		//err = application.ThrowError(application.TRANSACTION_ERROR, err.Error())
		return
	}
	defer func() {
		transactionContext.RollbackTransaction()
	}()

	var UserRepository, _ = repository.NewUserRepository(transactionContext) //factory.CreateUserRepository(transactionContext)
	var user *domain.Users
	if user, err = UserRepository.FindOne(common.ObjectToMap(request)); err != nil {
		//err = application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
		return
	}
	if user, err = UserRepository.Remove(user); err != nil {
		//err = application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
		return
	}
	rsp = user
	if err = transactionContext.CommitTransaction(); err != nil {
		//err = application.ThrowError(application.TRANSACTION_ERROR, err.Error())
		return
	}
	return
}

func (svr *UserService) ListUser(header *protocol.RequestHeader, request *protocolx.ListUserRequest) (rsp interface{}, err error) {
	var (
		transactionContext = transaction.NewPGTransactionContext(go_pg.DB)
		RoleRepository, _  = factory.CreateRoleRepository(transactionContext)
	)
	rsp = &protocolx.ListUserResponse{}
	if err = request.ValidateCommand(); err != nil {
		//err = application.ThrowError(application.TRANSACTION_ERROR, err.Error())
		return
	}
	if err = transactionContext.StartTransaction(); err != nil {
		log.Error(err)
		return nil, err
	}
	defer func() {
		transactionContext.RollbackTransaction()
	}()

	getRoles := func(roleIds []int64) string {
		if len(roleIds) == 0 {
			return ""
		}
		var roles []string
		var mapRoles = make(map[int64]*domain.Role)
		for _, id := range roleIds {
			if _, ok := mapRoles[id]; ok {
				continue
			}
			if roleItem, _ := RoleRepository.FindOne(map[string]interface{}{"id": id}); roleItem != nil {
				roles = append(roles, roleItem.RoleName)
			} else {
				mapRoles[id] = nil
			}
			continue
		}
		return strings.Join(roles, ",")
	}

	// TODO:可优化，每次只查询 user.id 列表 ，通过缓存查询user对象
	var UserRepository, _ = repository.NewUserRepository(transactionContext) //factory.CreateUserRepository(transactionContext)
	var user []*domain.Users
	var total int64
	if total, user, err = UserRepository.Find(common.ObjectToMap(request)); err != nil {
		//err = application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
		return
	}
	userList := make([]map[string]interface{}, 0)
	for _, v := range user {
		item := map[string]interface{}{"id": v.Id, "name": v.Name, "phone": v.Phone, "adminType": v.AdminType, "status": v.Status, "createTime": v.CreateTime.Local().Format("2006-01-02 15:04:05")} //v.CreateTime.Local().Format("2006-01-02 15:04:05")
		item["roles"] = getRoles(v.Roles)
		userList = append(userList, item)
	}

	rsp = map[string]interface{}{
		"total": total,
		//"pageNumber": request.PageNumber,
		"users": userList,
	}
	if err = transactionContext.CommitTransaction(); err != nil {
		//err = application.ThrowError(application.TRANSACTION_ERROR, err.Error())
		return
	}
	return
}

func NewUserService(options map[string]interface{}) *UserService {
	svr := &UserService{}
	return svr
}
