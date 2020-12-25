package role

import (
	"fmt"
	"github.com/astaxie/beego/validation"
)

type ListRoleResponse struct {
}

func (ListRoleResponse *ListRoleResponse) ValidateCommand() error {
	valid := validation.Validation{}
	b, err := valid.Valid(ListRoleResponse)
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
