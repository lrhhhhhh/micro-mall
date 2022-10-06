package activity

import (
	"context"
	"gateway/service/activity"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateActivityLogic {
	return &CreateActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateActivityLogic) CreateActivity(req *types.CreateActivityReq) (resp *types.CreateActivityResp, err error) {
	resp = &types.CreateActivityResp{}
	logx.Info("lrh pass")
	out, err := l.svcCtx.ActivityRpc.CreateActivity(l.ctx, &activity.ActivityReq{
		Id:             0,
		ActivityName:   req.ActivityName,
		GoodsId:        req.GoodsId,
		StockId:        req.StockId,
		StartTime:      req.StartTime,
		EndTime:        req.EndTime,
		Total:          req.Total,
		Status:         req.Status,
		BuyLimit:       req.BuyLimit,
		BuyProbability: req.BuyProbability,
	})
	if err != nil {
		resp.Message = err.Error()
	}

	resp.ActivityId = out.ActivityId
	resp.Message = out.Message

	return
}
