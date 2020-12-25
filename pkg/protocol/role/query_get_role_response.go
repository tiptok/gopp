package role

import (
	"fmt"
	"github.com/astaxie/beego/validation"
)

type GetRoleResponse struct {
}

func (GetRoleResponse *GetRoleResponse) ValidateCommand() error {
	valid := validation.Validation{}
	b, err := valid.Valid(GetRoleResponse)
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
