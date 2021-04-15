package user

import (
	"fmt"
	"github.com/beego/beego/v2/core/validation"
)

type DeleteUserResponse struct {
}

func (DeleteUserResponse *DeleteUserResponse) ValidateCommand() error {
	valid := validation.Validation{}
	b, err := valid.Valid(DeleteUserResponse)
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
