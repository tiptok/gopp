package models

import "time"

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
