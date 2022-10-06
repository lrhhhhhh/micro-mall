package stock

import (
	"context"
	"gateway/service/stock"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStockLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStockLogic {
	return &CreateStockLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateStockLogic) CreateStock(req *types.CreateStockReq) (resp *types.CreateStockResp, err error) {
	resp = &types.CreateStockResp{}
	out, err := l.svcCtx.StockRpc.CreateStock(l.ctx, &stock.CreateStockReq{GoodsId: req.GoodsId, Count: req.Count})
	if err != nil {
		resp.Message = err.Error()
	}

	logx.Info(out, err)

	resp.Message = out.Message
	resp.StockId = out.StockId

	return
}
