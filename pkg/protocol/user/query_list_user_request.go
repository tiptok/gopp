package user

import (
	"fmt"
	"github.com/astaxie/beego/validation"
)

type ListUserRequest struct {
	SearchByText string `json:"searchByText,omitempty"  form:"searchByText,optional"` //按名称搜索
	PageNumber   int    `json:"pageNumber,omitempty" valid:"Required" form:"pageNumber,optional"`
	Offset       int    `json:"offset,omitempty"  form:"offset,optional"`
	Limit        int    `json:"limit,omitempty"  form:"limit,optional"`
	SortById     string `json:"sortById,omitempty"`
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
