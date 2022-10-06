package stock

import (
	"context"
	"gateway/service/stock"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStockLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStockLogic {
	return &UpdateStockLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStockLogic) UpdateStock(req *types.UpdateStockReq) (resp *types.UpdateStockResp, err error) {
	resp = &types.UpdateStockResp{}
	out, err := l.svcCtx.StockRpc.UpdateStock(l.ctx, &stock.UpdateStockReq{
		StockId: req.StockId, GoodsId: req.GoodsId, Count: req.Count,
	})
	if err != nil {
		resp.Message = err.Error()
	}

	resp.Message = out.Message

	return
}
