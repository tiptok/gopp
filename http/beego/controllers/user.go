package controllers

import (
	"github.com/astaxie/beego/context"
	"github.com/tiptok/gocomm/pkg/log"
	"github.com/tiptok/gocomm/pkg/mybeego"
	service "github.com/tiptok/gopp/pkg/application/user"
	"github.com/tiptok/gopp/pkg/protocol"
	command "github.com/tiptok/gopp/pkg/protocol/user"
	"strconv"
)

type UserController struct {
	Base
	mybeego.ContextController
}

// CreateUser
// CreateUser execute command  create  User
func (controller *UserController) CreateUser(ctx *context.Context) {
	var (
		msg     *protocol.ResponseMessage
		svr     = service.NewUserService(nil)
		request *command.CreateUserRequest
	)
	defer func() {
		controller.Resp(ctx, msg)
	}()
	if err := controller.JsonUnmarshal(ctx, &request); err != nil {
		msg = protocol.NewResponseMessage(2, "")
		return
	}
	header := controller.GetRequestHeader(ctx)
	data, err := svr.CreateUser(header, request)
	if err != nil {
		log.Error(err)
	}
	msg = protocol.NewResponseMessageData(data, err)
}

// UpdateUser
// UpdateUser execute command  update  User
func (controller *UserController) UpdateUser(ctx *context.Context) {
	var (
		msg     *protocol.ResponseMessage
		svr     = service.NewUserService(nil)
		request *command.UpdateUserRequest
	)
	defer func() {
		controller.Resp(ctx, msg)
	}()
	if err := controller.JsonUnmarshal(ctx, &request); err != nil {
		msg = protocol.NewResponseMessage(2, "")
		return
	}
	request.Id, _ = strconv.ParseInt(ctx.Input.Query(":userId"), 10, 64)
	header := controller.GetRequestHeader(ctx)
	data, err := svr.UpdateUser(header, request)
	if err != nil {
		log.Error(err)
	}
	msg = protocol.NewResponseMessageData(data, err)
}

// GetUser
// GetUser execute query  get  User
func (controller *UserController) GetUser(ctx *context.Context) {
	var (
		msg     *protocol.ResponseMessage
		svr     = service.NewUserService(nil)
		request = &command.GetUserRequest{}
	)
	defer func() {
		controller.Resp(ctx, msg)
	}()
	request.Id, _ = strconv.ParseInt(ctx.Input.Query(":userId"), 10, 64)
	header := controller.GetRequestHeader(ctx)
	data, err := svr.GetUser(header, request)
	if err != nil {
		log.Error(err)
	}
	msg = protocol.NewResponseMessageData(data, err)
}

// DeleteUser
// DeleteUser execute command  delete  User
func (controller *UserController) DeleteUser(ctx *context.Context) {
	var (
		msg     *protocol.ResponseMessage
		svr     = service.NewUserService(nil)
		request = &command.DeleteUserRequest{}
	)
	defer func() {
		controller.Resp(ctx, msg)
	}()
	request.Id, _ = strconv.ParseInt(ctx.Input.Query(":userId"), 10, 64)
	header := controller.GetRequestHeader(ctx)
	data, err := svr.DeleteUser(header, request)
	if err != nil {
		log.Error(err)
	}
	msg = protocol.NewResponseMessageData(data, err)
}

// ListUser
// ListUser execute query  list  User
func (controller *UserController) ListUser(ctx *context.Context) {
	var (
		msg     *protocol.ResponseMessage
		svr     = service.NewUserService(nil)
		request = &command.ListUserRequest{}
	)
	defer func() {
		controller.Resp(ctx, msg)
	}()
	request.PageNumber, _ = strconv.Atoi(ctx.Input.Query("pageNumber"))

	request.Offset, request.Limit = controller.GetLimitInfo(ctx)
	request.SearchByText = ctx.Input.Query("searchByText")
	header := controller.GetRequestHeader(ctx)
	data, err := svr.ListUser(header, request)
	if err != nil {
		log.Error(err)
	}
	msg = protocol.NewResponseMessageListData(data, err)
}
