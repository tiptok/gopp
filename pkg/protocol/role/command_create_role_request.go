package role

import (
	"fmt"
	"github.com/beego/beego/v2/core/validation"
)

type CreateRoleRequest struct {
	// 角色名称
	RoleName string `json:"roleName,omitempty"`
	// 父级Id
	ParentId int64 `json:"parentId,omitempty"`
}

func (CreateRoleRequest *CreateRoleRequest) ValidateCommand() error {
	valid := validation.Validation{}
	b, err := valid.Valid(CreateRoleRequest)
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
