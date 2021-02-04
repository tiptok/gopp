package protocol

import (
	"encoding/json"
)

//CustomErrParse 解析自定义错误结构体
type CustomErrParse interface {
	ParseToMessage() *ResponseMessage
}

//ErrorMap 统一消息错误编码
type ErrorMap map[int]string

//Search 搜索错误描述
func (m ErrorMap) Search(code int) ErrorCode {
	if v, ok := m[code]; ok {
		return ErrorCode{
			Errno:  code,
			Errmsg: v,
		}
	}
	return ErrorCode{Errno: code, Errmsg: "错误码未定义"}
}

func NewMesage(code int) *ResponseMessage {
	return &ResponseMessage{
		ErrorCode: SearchErr(code),
		Data: struct {
		}{},
	}
}

var (
	_ CustomErrParse = new(ErrWithMessage)
	_ error          = new(ErrWithMessage)
)

//Error 实现接口error 中的方法
//将ErrorCode转为json数据，建议用于日志记录
func (e ErrWithMessage) Error() string {
	bt, _ := json.Marshal(e.ErrorCode)
	return string(bt)
}

//Unwrap 接口实现
func (e ErrWithMessage) Unwrap() error {
	return e.Err
}

//ParseToMessage 实现CustomErrParse的接口
func (e ErrWithMessage) ParseToMessage() *ResponseMessage {
	return &ResponseMessage{
		ErrorCode: e.ErrorCode,
		Data:      nil,
	}
}

func SearchErr(code int) ErrorCode {
	return errmessge.Search(code)
}
func NewResponseMessageData(data interface{}, err error) *ResponseMessage {
	var msg *ResponseMessage
	if err == nil {
		msg = NewMesage(0)
		msg.Data = data
		return msg
	}
	//log.Error("服务错误:" + eRR.Error())
	if x, ok := err.(CustomErrParse); ok {
		msg = x.ParseToMessage()
		msg.Data = data
		return msg
	}
	//if v, ok := err.(*application.ServiceError); ok {
	//	msg = &ResponseMessage{
	//		ErrorCode{v.Code, v.Message},
	//		data,
	//	}
	//	return msg
	//}
	return NewMesage(1)
}

func NewResponseMessageListData(data interface{}, err error) *ResponseMessage {
	mapData := map[string]interface{}{"gridResult": data}
	return NewResponseMessageData(mapData, err)
}

func NewResponseMessage(code int, err string) *ResponseMessage {
	return &ResponseMessage{
		ErrorCode: ErrorCode{
			Errno:  code,
			Errmsg: err,
		},
		Data: struct {
		}{},
	}
}

func NewCustomMessage(code int, msg string) *ErrWithMessage {
	return &ErrWithMessage{
		Err:       nil,
		ErrorCode: ErrorCode{code, msg},
	}
}

//ErrorCode 统一错误结构
type ErrorCode struct {
	Errno  int    `json:"code"`
	Errmsg string `json:"msg"`
}

//ResponseMessage 统一返回消息结构体
type ResponseMessage struct {
	ErrorCode
	Data interface{} `json:"data"`
}

//ErrWithMessage  自定义错误结构
type ErrWithMessage struct {
	Err error `json:"-"`
	ErrorCode
}

var errmessge ErrorMap = map[int]string{
	0: "成功",
	1: "系统异常",
	2: "参数错误",
}

type RequestHeader struct {
	UserId   int64 //UserId 唯一标识
	UserName string
	Token    string
	BodyKeys []string
}
