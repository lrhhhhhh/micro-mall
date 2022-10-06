package seckill

import (
	"context"
	"gateway/service/seckill"
	"time"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Seckill2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSeckill2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Seckill2Logic {
	return &Seckill2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Seckill2Logic) Seckill2(req *types.SeckillReq) (resp *types.SeckillResp, err error) {
	res, err := l.svcCtx.SeckillRpc.Seckill2(l.ctx, &seckill.SeckillReq{
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
