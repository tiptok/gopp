package user

import (
	"fmt"
	"github.com/astaxie/beego/validation"
)

type GetUserRequest struct {
	// 唯一标识
	Id int64 `json:"id,omitempty"`
}

func (GetUserRequest *GetUserRequest) ValidateCommand() error {
	valid := validation.Validation{}
	b, err := valid.Valid(GetUserRequest)
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
