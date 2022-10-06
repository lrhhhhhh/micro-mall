package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"activity/internal/svc"
	"activity/service/activity"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityLogic {
	return &GetActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetActivityLogic) GetActivity(in *activity.ActivityReq) (*activity.ActivityResp, error) {
	record, err := l.svcCtx.ActivityModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &activity.ActivityResp{Activity: &activity.ActivityModel{
		Id:             int64(record.Id),
		ActivityName:   record.ActivityName,
		GoodsId:        record.GoodsId,
		StockId:        record.StockId,
		StartTime:      record.StartTime,
		EndTime:        record.EndTime,
		Total:          record.Total,
		Status:         record.Status,
		BuyLimit:       record.BuyLimit,
		BuyProbability: record.BuyProbability,
	}}, nil
}
