package domain

import (
	"fmt"
	"github.com/tiptok/gopp/pkg/constant"
	"time"
)

const (
	UserAdmin = iota + 1
	UserNormal
)

// Users
type Users struct {
	// 唯一标识
	Id int64 `json:"id"`
	// 名称
	Name string `json:"name"`
	// 手机号
	Phone string `json:"phone"`
	// 密码
	Passwd string `json:"-"`
	// 用户角色
	Roles []int64 `json:"roles"`
	// 1启用  2禁用
	Status int `json:"status"`
	// 管理员类型 1:管理员  2：普通员工
	AdminType int `json:"adminType"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
}

type UsersRepository interface {
	Save(dm *Users) (*Users, error)
	Remove(dm *Users) (*Users, error)
	FindOne(queryOptions map[string]interface{}) (*Users, error)
	FindOneByPhone(phone string) (*Users, error)
	FindOneByPhoneNoCache(phone string) (*Users, error)
	Find(queryOptions map[string]interface{}) (int64, []*Users, error)
}

func (m *Users) Identify() interface{} {
	if m.Id == 0 {
		return nil
	}
	return m.Id
}

func (m *Users) CacheKey() string {
	return fmt.Sprintf("%v:cache:users:id:%v", constant.POSTGRESQL_DB_NAME, m.Id)
}

func (m *Users) Update(data map[string]interface{}) error {
	if name, ok := data["name"]; ok {
		m.Name = name.(string)
	}
	if phone, ok := data["phone"]; ok {
		m.Phone = phone.(string)
	}
	if Passwd, ok := data["passwd"]; ok {
		m.Passwd = Passwd.(string)
	}
	if roles, ok := data["roles"]; ok {
		m.Roles = roles.([]int64)
	}
	if adminType, ok := data["adminType"]; ok {
		m.AdminType = adminType.(int)
	}
	if status, ok := data["status"]; ok {
		m.Status = status.(int)
	}
	m.UpdateTime = time.Now()
	return nil
}
