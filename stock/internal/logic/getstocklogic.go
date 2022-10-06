package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"stock/internal/svc"
	"stock/service/stock"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStockLogic {
	return &GetStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStockLogic) GetStock(in *stock.GetStockReq) (*stock.GetStockResp, error) {
	res, err := l.svcCtx.StockModel.FindOne(l.ctx, in.StockId)

	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &stock.GetStockResp{
		Message: "success",
		Stock: &stock.GetStockResp_StockModel{
			StockId: res.Id,
			GoodsId: res.GoodsId,
			Count:   res.Count,
		},
	}, nil
}
