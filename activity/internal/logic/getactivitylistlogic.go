package logic

import (
	"context"

	"activity/internal/svc"
	"activity/service/activity"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActivityListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetActivityListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityListLogic {
	return &GetActivityListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetActivityListLogic) GetActivityList(in *activity.ActivityListReq) (*activity.ActivityListResp, error) {
	// todo: add your logic here and delete this line

	return &activity.ActivityListResp{}, nil
}
