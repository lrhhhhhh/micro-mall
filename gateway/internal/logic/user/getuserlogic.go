package user

import (
	"context"
	"gateway/internal/svc"
	"gateway/internal/types"
	"gateway/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req types.GetUserReq) (resp *types.GetUserResp, err error) {
	u, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.GetUserReq{Uid: req.Uid})
	if err != nil {
		return nil, err
	}

	resp = &types.GetUserResp{}
	resp.User.Uid = u.User.ID
	resp.User.Username = u.User.UserName
	resp.Message = "success"

	return resp, nil
}
