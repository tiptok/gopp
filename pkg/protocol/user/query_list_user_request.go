package user

import (
	"fmt"
	"github.com/astaxie/beego/validation"
)

type ListUserRequest struct {
	SearchByText string `json:"searchByText,omitempty"` //按名称搜索
	//PageSize     int    `json:"pageSize" valid:"Required"`
	PageNumber int `json:"pageNumber" valid:"Required"`

	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	SortById string `json:"sortById"`
}

func (ListUserRequest *ListUserRequest) ValidateCommand() error {
	valid := validation.Validation{}
	//if ListUserRequest.PageSize == 0 {
	//	ListUserRequest.PageSize = 20
	//}
	if len(ListUserRequest.SortById) == 0 {
		ListUserRequest.SortById = "DESC"
	}
	b, err := valid.Valid(ListUserRequest)
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
