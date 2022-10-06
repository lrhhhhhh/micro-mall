package stock

import (
	"context"
	"gateway/service/stock"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStockLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStockLogic {
	return &GetStockLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStockLogic) GetStock(req *types.GetStockReq) (resp *types.GetStockResp, err error) {
	resp = &types.GetStockResp{}
	out, err := l.svcCtx.StockRpc.GetStock(l.ctx, &stock.GetStockReq{StockId: req.StockId})
	if err != nil {
		resp.Message = err.Error()
		return
	}

	resp.Message = out.Message
	resp.Stock = types.Stock{
		StockId: out.Stock.StockId,
		GoodsId: out.Stock.GoodsId,
		Count:   out.Stock.Count,
	}

	return
}
