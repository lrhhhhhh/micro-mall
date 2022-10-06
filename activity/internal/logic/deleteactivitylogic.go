package logic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"activity/internal/svc"
	"activity/service/activity"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteActivityLogic {
	return &DeleteActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteActivityLogic) DeleteActivity(in *activity.ActivityReq) (*activity.BaseActivityResp, error) {
	err := l.svcCtx.ActivityModel.Delete(l.ctx, uint64(in.Id))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	_, err = l.svcCtx.Redis.Del(
		fmt.Sprintf("SeckillActivity:%d", in.Id),
		//fmt.Sprintf("SeckillHistory:%d", in.Id),
		fmt.Sprintf(l.svcCtx.Config.StockRedisKeyFormat, in.Id, in.GoodsId, in.StockId),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &activity.BaseActivityResp{}, err
}
