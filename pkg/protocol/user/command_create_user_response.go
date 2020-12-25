package user

import (
	"fmt"
	"github.com/astaxie/beego/validation"
)

type CreateUserResponse struct {
}

func (CreateUserResponse *CreateUserResponse) ValidateCommand() error {
	valid := validation.Validation{}
	b, err := valid.Valid(CreateUserResponse)
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
