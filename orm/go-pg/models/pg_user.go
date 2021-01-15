package models

import "time"

// Users
type Users struct {
	tableName struct{} `pg:"users"`
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
	AdminType int
	// 创建时间
	CreateTime time.Time
	// 更新时间
	UpdateTime time.Time
}
