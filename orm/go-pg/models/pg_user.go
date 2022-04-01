package models

import (
	"fmt"
	"time"
)

// Users
type Users struct {
	tableName struct{} `pg:"users"` //指定schema `pg:"base.users"`
	//	唯一标识
	Id int64
	//	名称
	Name string
	//	手机号
	Phone string
	//	密码
	Passwd string
	//	用户角色
	Roles []int64
	// 1启用  2禁用
	Status int
	// 管理员类型 1:超级管理员  2：普通账号
	AdminType int `pg:"default:2"`
	// 创建时间
	CreateTime time.Time `pg:"default:current_timestamp"`
	// 更新时间
	UpdateTime time.Time
}

func (m *Users) CacheKeyFunc() string {
	if m.Id == 0 {
		return ""
	}
	return fmt.Sprintf("%v:cache:users:id:%v", "gopp", m.Id)
}

func (m *Users) CachePrimaryKeyFunc() string {
	if len(m.Phone) == 0 {
		return ""
	}
	return fmt.Sprintf("%v:cache:users:phone:%v", "gopp", m.Phone)
}
