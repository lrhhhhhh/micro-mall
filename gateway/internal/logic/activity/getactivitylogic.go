package activity

import (
	"context"
	"gateway/service/activity"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityLogic {
	return &GetActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetActivityLogic) GetActivity(req *types.GetActivityReq) (resp *types.GetActivityResp, err error) {
	resp = &types.GetActivityResp{}
	out, err := l.svcCtx.ActivityRpc.GetActivity(l.ctx, &activity.ActivityReq{Id: req.ActivityId})
	if err != nil {
		resp.Message = err.Error()
	}

	resp.Message = "success"
	resp.Activity = types.Activity{
		Id:             out.Activity.Id,
		ActivityName:   out.Activity.ActivityName,
		GoodsId:        out.Activity.GoodsId,
		StockId:        out.Activity.StockId,
		StartTime:      out.Activity.StartTime,
		EndTime:        out.Activity.EndTime,
		Total:          out.Activity.Total,
		Status:         out.Activity.Status,
		BuyLimit:       out.Activity.BuyLimit,
		BuyProbability: out.Activity.BuyProbability,
	}

	return
}
