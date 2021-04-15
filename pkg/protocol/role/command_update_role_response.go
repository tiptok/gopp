package role

import (
	"fmt"
	"github.com/beego/beego/v2/core/validation"
)

type UpdateRoleResponse struct {
}

func (UpdateRoleResponse *UpdateRoleResponse) ValidateCommand() error {
	valid := validation.Validation{}
	b, err := valid.Valid(UpdateRoleResponse)
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
