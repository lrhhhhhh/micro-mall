package seckill

import (
	"context"
	"gateway/service/seckill"
	"time"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SeckillLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSeckillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillLogic {
	return &SeckillLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SeckillLogic) Seckill(req *types.SeckillReq) (resp *types.SeckillResp, err error) {
	res, err := l.svcCtx.SeckillRpc.Seckill(l.ctx, &seckill.SeckillReq{
		Uid:        req.Uid,
		ActivityId: req.ActivityId,
		GoodsId:    req.GoodsId,
		StockId:    req.StockId,
		BuyCnt:     req.BuyCnt,
		AccessTime: time.Now().Unix(),
	})
	if err != nil {
		return nil, err
	}

	return &types.SeckillResp{
		Code:    res.Code,
		Message: res.Message,
	}, nil
}
