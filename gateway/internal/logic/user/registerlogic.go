package user

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"
	"gateway/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.BaseResp, err error) {
	resp = &types.BaseResp{}
	res, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterReq{UserName: req.UserName, Password: req.Password, PasswordConfirm: req.PasswordConfirm})
	resp.Message = res.Message
	if err != nil {
		return resp, err
	}

	return resp, nil
}
