package activity

import (
	"context"
	"gateway/service/activity"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateActivityLogic {
	return &UpdateActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateActivityLogic) UpdateActivity(req *types.UpdateActivityReq) (resp *types.UpdateActivityResp, err error) {
	resp = &types.UpdateActivityResp{}
	out, err := l.svcCtx.ActivityRpc.UpdateActivity(l.ctx, &activity.ActivityReq{
		Id:             req.Id,
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
		return
	}

	resp.Message = out.Message

	return
}
