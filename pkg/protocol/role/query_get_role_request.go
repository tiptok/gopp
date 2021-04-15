package role

import (
	"fmt"
	"github.com/beego/beego/v2/core/validation"
)

type GetRoleRequest struct {
	// 唯一标识
	Id int64 `json:"id" valid:"Required"`
}

func (GetRoleRequest *GetRoleRequest) ValidateCommand() error {
	valid := validation.Validation{}
	b, err := valid.Valid(GetRoleRequest)
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
