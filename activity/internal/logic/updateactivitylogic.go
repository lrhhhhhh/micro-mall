package logic

import (
	"activity/internal/model"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"

	"activity/internal/svc"
	"activity/service/activity"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateActivityLogic {
	return &UpdateActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateActivityLogic) UpdateActivity(in *activity.ActivityReq) (*activity.BaseActivityResp, error) {
	err := l.svcCtx.ActivityModel.Update(l.ctx, &model.Activity{
		Id:             uint64(in.Id),
		ActivityName:   in.ActivityName,
		GoodsId:        in.GoodsId,
		StockId:        in.StockId,
		StartTime:      in.StartTime,
		EndTime:        in.EndTime,
		Total:          in.Total,
		Status:         in.Status,
		BuyLimit:       in.BuyLimit,
		BuyProbability: in.BuyProbability,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	key := fmt.Sprintf("SeckillActivity:%d", in.Id)
	err = l.svcCtx.Redis.Hmset(key, map[string]string{
		"ActivityId":     strconv.FormatInt(in.Id, 10),
		"ActivityName":   in.ActivityName,
		"GoodsId":        strconv.FormatInt(in.GoodsId, 10),
		"StartTime":      strconv.FormatInt(in.StartTime, 10),
		"EndTime":        strconv.FormatInt(in.EndTime, 10),
		"Total":          strconv.FormatInt(in.Total, 10),
		"Status":         strconv.FormatInt(in.Status, 10),
		"BuyLimit":       strconv.FormatInt(in.BuyLimit, 10),
		"BuyProbability": strconv.FormatFloat(in.BuyProbability, 'f', 10, 64),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = l.svcCtx.Redis.SetCtx(
		l.ctx,
		fmt.Sprintf(l.svcCtx.Config.StockRedisKeyFormat, in.Id, in.GoodsId, in.StockId),
		strconv.FormatInt(in.Total, 10),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &activity.BaseActivityResp{Message: "更新成功"}, nil
}
