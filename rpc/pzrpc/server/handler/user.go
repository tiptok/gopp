package handler

import (
	"context"
	"github.com/tiptok/gopp/rpc/pzrpc/protobuf"
)

type UserServer struct {
}

func NewUserServer() *UserServer {
	return &UserServer{}
}
func (*UserServer) GetUser(context.Context, *protobuf.GetUsersReq) (*protobuf.GetUsersResp, error) {
	return &protobuf.GetUsersResp{
		User: &protobuf.User{
			Name:  "tip",
			Phone: "180000000001",
			Roles: []int64{1, 2},
		},
	}, nil
}
func (*UserServer) CreateUser(context.Context, *protobuf.CreateUserReq) (*protobuf.CreateUserResp, error) {
	return nil, nil
}
func (*UserServer) UpdateUser(context.Context, *protobuf.UpdateUserReq) (*protobuf.UpdateUserResp, error) {
	return nil, nil
}
func (*UserServer) RemoveUser(context.Context, *protobuf.RemoveUsersReq) (*protobuf.RemoveUsersResp, error) {
	return nil, nil
}
func (*UserServer) ListUser(context.Context, *protobuf.ListUserReq) (*protobuf.ListUserResp, error) {
	// 实现逻辑
	return nil, nil
}
