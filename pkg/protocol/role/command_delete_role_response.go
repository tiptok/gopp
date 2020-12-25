package role

import (
	"fmt"
	"github.com/astaxie/beego/validation"
)

type DeleteRoleResponse struct {
}

func (DeleteRoleResponse *DeleteRoleResponse) ValidateCommand() error {
	valid := validation.Validation{}
	b, err := valid.Valid(DeleteRoleResponse)
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
