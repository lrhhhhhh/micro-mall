// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package userclient

import (
	"context"

	"user/service/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetUserReq   = user.GetUserReq
	LoginReq     = user.LoginReq
	RegisterReq  = user.RegisterReq
	UserModel    = user.UserModel
	UserResponse = user.UserResponse

	User interface {
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*UserResponse, error)
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*UserResponse, error)
		GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*UserResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*UserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUser) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*UserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUser) GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*UserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}
