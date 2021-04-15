package role

import (
	"fmt"
	"github.com/beego/beego/v2/core/validation"
)

type ListRoleRequest struct {
}

func (ListRoleRequest *ListRoleRequest) ValidateCommand() error {
	valid := validation.Validation{}
	b, err := valid.Valid(ListRoleRequest)
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
