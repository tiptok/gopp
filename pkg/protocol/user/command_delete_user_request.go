package user

import (
	"fmt"
	"github.com/beego/beego/v2/core/validation"
)

type DeleteUserRequest struct {
	// 唯一标识
	Id int64 `json:"id,omitempty" path:"userId"`
}

func (DeleteUserRequest *DeleteUserRequest) ValidateCommand() error {
	valid := validation.Validation{}
	b, err := valid.Valid(DeleteUserRequest)
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
