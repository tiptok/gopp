package handler

import (
	"context"
	"github.com/tiptok/gopp/rpc/pzrpc/protocol"
)

type UserServer struct {
}

func NewUserServer() *UserServer {
	return &UserServer{}
}
func (*UserServer) GetUser(context.Context, *protocol.GetUsersReq) (*protocol.GetUsersResp, error) {
	return &protocol.GetUsersResp{
		User: &protocol.User{
			Name:  "tip",
			Phone: "180000000001",
			Roles: []int64{1, 2},
		},
	}, nil
}
func (*UserServer) CreateUser(context.Context, *protocol.CreateUserReq) (*protocol.CreateUserResp, error) {
	return nil, nil
}
func (*UserServer) UpdateUser(context.Context, *protocol.UpdateUserReq) (*protocol.UpdateUserResp, error) {
	return nil, nil
}
func (*UserServer) RemoveUser(context.Context, *protocol.RemoveUsersReq) (*protocol.RemoveUsersResp, error) {
	return nil, nil
}
func (*UserServer) ListUser(context.Context, *protocol.ListUserReq) (*protocol.ListUserResp, error) {
	// 实现逻辑
	return nil, nil
}
