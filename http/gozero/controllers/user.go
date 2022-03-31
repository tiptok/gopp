package controllers

import (
	"github.com/tiptok/gocomm/pkg/log"
	service "github.com/tiptok/gopp/pkg/application/user"
	"github.com/tiptok/gopp/pkg/protocol"
	command "github.com/tiptok/gopp/pkg/protocol/user"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type UserController struct {
	Base
}

// CreateUser
// CreateUser execute command  create  User
func (controller *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var (
		msg     *protocol.ResponseMessage
		svr     = service.NewUserService(nil)
		request *command.CreateUserRequest
	)
	defer func() {
		controller.Resp(w, msg)
	}()
	if err := controller.JsonUnmarshal(r, &request); err != nil {
		msg = protocol.NewResponseMessage(2, err.Error())
		return
	}
	header := controller.GetRequestHeader(r)
	data, err := svr.CreateUser(header, request)
	if err != nil {
		log.Error(err)
	}
	msg = protocol.NewResponseMessageData(data, err)
}

// UpdateUser
// UpdateUser execute command  update  User
func (controller *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var (
		msg     *protocol.ResponseMessage
		svr     = service.NewUserService(nil)
		request *command.UpdateUserRequest
	)
	defer func() {
		controller.Resp(w, msg)
	}()
	if err := controller.JsonUnmarshal(r, &request); err != nil {
		msg = protocol.NewResponseMessage(2, err.Error())
		return
	}
	if err := httpx.ParsePath(r, request); err != nil {
		msg = protocol.NewResponseMessage(2, err.Error())
		return
	}
	header := controller.GetRequestHeader(r)
	data, err := svr.UpdateUser(header, request)
	if err != nil {
		log.Error(err)
	}
	msg = protocol.NewResponseMessageData(data, err)
}

// GetUser
// GetUser execute query  get  User
func (controller *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	var (
		msg     *protocol.ResponseMessage
		svr     = service.NewUserService(nil)
		request = &command.GetUserRequest{}
	)
	defer func() {
		controller.Resp(w, msg)
	}()
	if err := httpx.ParsePath(r, request); err != nil {
		msg = protocol.NewResponseMessage(2, err.Error())
		return
	}
	header := controller.GetRequestHeader(r)
	data, err := svr.GetUser(header, request)
	if err != nil {
		log.Error(err)
	}
	msg = protocol.NewResponseMessageData(data, err)
}

// DeleteUser
// DeleteUser execute command  delete  User
func (controller *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var (
		msg     *protocol.ResponseMessage
		svr     = service.NewUserService(nil)
		request = &command.DeleteUserRequest{}
	)
	defer func() {
		controller.Resp(w, msg)
	}()
	if err := httpx.ParsePath(r, request); err != nil {
		msg = protocol.NewResponseMessage(2, err.Error())
		return
	}
	header := controller.GetRequestHeader(r)
	data, err := svr.DeleteUser(header, request)
	if err != nil {
		log.Error(err)
	}
	msg = protocol.NewResponseMessageData(data, err)
}

// ListUser
// ListUser execute query  list  User
func (controller *UserController) ListUser(w http.ResponseWriter, r *http.Request) {
	var (
		msg     *protocol.ResponseMessage
		svr     = service.NewUserService(nil)
		request = &command.ListUserRequest{}
	)
	defer func() {
		controller.Resp(w, msg)
	}()
	if err := httpx.ParseForm(r, request); err != nil {
		msg = protocol.NewResponseMessage(2, err.Error())
		return
	}
	header := controller.GetRequestHeader(r)
	data, err := svr.ListUser(header, request)
	if err != nil {
		log.Error(err)
	}
	msg = protocol.NewResponseMessageListData(data, err)
}
