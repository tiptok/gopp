package user

import (
	"fmt"
	"github.com/astaxie/beego/validation"
)

type UpdateUserResponse struct {
}

func (UpdateUserResponse *UpdateUserResponse) ValidateCommand() error {
	valid := validation.Validation{}
	b, err := valid.Valid(UpdateUserResponse)
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
