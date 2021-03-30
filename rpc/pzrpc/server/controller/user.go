package controller

import (
	"context"
	service "github.com/tiptok/gopp/pkg/application/user"
	"github.com/tiptok/gopp/pkg/protobuf/user"
	command "github.com/tiptok/gopp/pkg/protocol/user"
)

type UserController struct {
	Base
}

func NewUserController() *UserController {
	return &UserController{}
}
func (controller *UserController) GetUser(ctx context.Context, req *user.GetUsersReq) (*user.GetUsersResp, error) {
	var (
		svr     = service.NewUserService(nil)
		request = &command.GetUserRequest{}
		resp    = &user.GetUsersResp{}
	)
	controller.JsonUnmarshal(request, req)
	response, err := svr.GetUser(nil, request)
	if err != nil {
		return nil, err
	}
	controller.JsonUnmarshal(resp, response)
	return resp, nil
}
func (controller *UserController) CreateUser(ctx context.Context, req *user.CreateUserReq) (*user.CreateUserResp, error) {
	var (
		svr     = service.NewUserService(nil)
		request = &command.CreateUserRequest{}
		resp    = &user.CreateUserResp{User: &user.User{}}
	)
	controller.JsonUnmarshal(request, req)
	response, err := svr.CreateUser(nil, request)
	if err != nil {
		return nil, err
	}
	if err = controller.JsonUnmarshal(resp.User, response); err != nil {
		return nil, err
	}
	return resp, nil
}
func (controller *UserController) UpdateUser(ctx context.Context, req *user.UpdateUserReq) (*user.UpdateUserResp, error) {
	var (
		svr     = service.NewUserService(nil)
		request = &command.UpdateUserRequest{}
		resp    = &user.UpdateUserResp{User: &user.User{}}
	)
	controller.JsonUnmarshal(request, req)
	header := controller.GetRequestHeader(ctx, req)
	response, err := svr.UpdateUser(header, request)
	if err != nil {
		return nil, err
	}
	controller.JsonUnmarshal(resp.User, response)
	return resp, nil
}
func (controller *UserController) RemoveUser(ctx context.Context, req *user.RemoveUsersReq) (*user.RemoveUsersResp, error) {
	var (
		svr     = service.NewUserService(nil)
		request = &command.DeleteUserRequest{}
		resp    = &user.RemoveUsersResp{User: &user.User{}}
	)
	controller.JsonUnmarshal(request, req)
	response, err := svr.DeleteUser(nil, request)
	if err != nil {
		return nil, err
	}
	controller.JsonUnmarshal(resp.User, response)
	return resp, nil
}
func (controller *UserController) ListUser(ctx context.Context, req *user.ListUserReq) (*user.ListUserResp, error) {
	var (
		svr     = service.NewUserService(nil)
		request = &command.ListUserRequest{}
		resp    = &user.ListUserResp{}
	)
	controller.JsonUnmarshal(request, req)
	response, err := svr.ListUser(nil, request)
	if err != nil {
		return nil, err
	}
	controller.JsonUnmarshal(resp, response)
	return resp, nil
}
