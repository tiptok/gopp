package models

import (
	"fmt"
	"time"
)

// Role
type Role struct {
	tableName struct{} `pg:"role"`
	//	唯一标识
	Id int64
	//	角色名称
	RoleName string
	//	父级Id
	ParentId int64
	//	创建时间
	CreateTime time.Time
	//	更新时间
	UpdateTime time.Time
}

func (m *Role) CacheKeyFunc() string {
	if m.Id == 0 {
		return ""
	}
	return fmt.Sprintf("%v:cache:role:id:%v", "gopp", m.Id)
}
