package logic

import (
	"context"
	"stock/internal/model"

	"stock/internal/svc"
	"stock/service/stock"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStockLogic {
	return &UpdateStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStockLogic) UpdateStock(in *stock.UpdateStockReq) (*stock.UpdateStockResp, error) {
	// todo: firstly, check goods_id existence

	if in.Count < 0 {
		return &stock.UpdateStockResp{Message: "数量不能为负数"}, nil
	}

	err := l.svcCtx.StockModel.Update(l.ctx, &model.Stock{
		Id:      in.StockId,
		GoodsId: in.GoodsId,
		Count:   in.Count,
	})
	if err != nil {
		return &stock.UpdateStockResp{Message: err.Error()}, nil
	}

	return &stock.UpdateStockResp{Message: "更新成功"}, nil
}
