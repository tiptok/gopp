package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Users struct {
	//	唯一标识
	Id int64 `gorm:"primaryKey"`
	//	名称
	Name string
	//	手机号
	Phone string
	//	密码
	Passwd string
	//	用户角色
	Roles SliceInt64 `gorm:"type:jsonb;not null;default '{}'"`
	// 1启用  2禁用
	Status int
	// 管理员类型 1:超级管理员  2：普通账号
	AdminType int
	// 创建时间
	CreateTime time.Time
	// 更新时间
	UpdateTime time.Time
	//gorm.Model
}

func (m *Users) TableName() string {
	return "users"
}

type SliceInt64 []int64

func (p SliceInt64) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}
func (p *SliceInt64) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	var i []int64
	if err := json.Unmarshal(source, &i); err != nil {
		return err
	}

	*p = i
	return nil
}
