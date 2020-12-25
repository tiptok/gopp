package role

import (
	"fmt"
	"github.com/astaxie/beego/validation"
)

type UpdateRoleRequest struct {
	// 角色Id
	Id int64 `json:"id,omitempty" valid:"Required"`
	// 角色名称
	RoleName string `json:"roleName,omitempty"`
	// 父级Id
	ParentId int64 `json:"parentId,omitempty"`
}

func (UpdateRoleRequest *UpdateRoleRequest) ValidateCommand() error {
	valid := validation.Validation{}
	b, err := valid.Valid(UpdateRoleRequest)
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
