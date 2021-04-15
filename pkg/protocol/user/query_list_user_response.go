package user

import (
	"fmt"
	"github.com/beego/beego/v2/core/validation"
)

type ListUserResponse struct {
}

func (ListUserResponse *ListUserResponse) ValidateCommand() error {
	valid := validation.Validation{}
	b, err := valid.Valid(ListUserResponse)
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
