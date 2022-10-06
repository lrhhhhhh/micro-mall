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

type CreateActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateActivityLogic {
	return &CreateActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateActivityLogic) CreateActivity(in *activity.ActivityReq) (*activity.CreateActivityResp, error) {
	res, err := l.svcCtx.ActivityModel.Insert(l.ctx, &model.Activity{
		Id:             0,
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

	lastId, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	key := fmt.Sprintf(l.svcCtx.Config.ActivityRedisKeyFormat, lastId)
	err = l.svcCtx.Redis.Hmset(key, map[string]string{
		"ActivityId":     strconv.FormatInt(lastId, 10),
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

	// 创建history哈希表，用一个magic number占位
	_, err = l.svcCtx.Redis.Hsetnx(fmt.Sprintf(l.svcCtx.Config.HistoryRedisKeyFormat, lastId), "-0x3f3f3f3f", "")
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	_, err = l.svcCtx.Redis.Setnx(
		fmt.Sprintf(l.svcCtx.Config.StockRedisKeyFormat, lastId, in.GoodsId, in.StockId),
		strconv.FormatInt(in.Total, 10),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &activity.CreateActivityResp{ActivityId: lastId, Message: "创建成功"}, nil
}
