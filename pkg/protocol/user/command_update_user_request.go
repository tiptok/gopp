package user

import (
	"fmt"
	"github.com/astaxie/beego/validation"
)

type UpdateUserRequest struct {
	Id int64 `json:"-" valid:"Required" path:"userId"`
	// 名称
	Name string `json:"name,omitempty"`
	// 地址
	Address string `json:"address,omitempty"`
	// 手机号
	Phone string `json:"phone,omitempty"`
	// 密码
	Passwd string `json:"passwd,omitempty"`
	// 用户角色
	Roles []int64 `json:"roles,omitempty"`
	// 1启用  2禁用
	Status int `json:"status,omitempty"`
	// 管理员类型  1:超级管理员  2：普通账号
	AdminType int `json:"adminType"`
}

func (UpdateUserRequest *UpdateUserRequest) ValidateCommand() error {
	valid := validation.Validation{}
	b, err := valid.Valid(UpdateUserRequest)
	if err != nil {
		return err
	}
	if !b {
		for _, validErr := range valid.Errors {
			return fmt.Errorf("%s  %s", validErr.Key, validErr.Message)
		}
	}
	return nil
}
