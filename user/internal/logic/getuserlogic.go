package logic

import (
	"context"
	"user/internal/svc"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.GetUserReq) (*user.UserResponse, error) {
	u, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Uid)
	if err != nil {
		return nil, err
	}
	return &user.UserResponse{User: model2Pb(u), Message: "success"}, nil
}
