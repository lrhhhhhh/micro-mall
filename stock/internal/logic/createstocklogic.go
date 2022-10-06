package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"stock/internal/model"

	"stock/internal/svc"
	"stock/service/stock"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStockLogic {
	return &CreateStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateStockLogic) CreateStock(in *stock.CreateStockReq) (*stock.CreateStockResp, error) {
	// todo: firstly, check goods_id existence

	if in.Count < 0 {
		return &stock.CreateStockResp{Message: "数量不能为负数"}, nil
	}

	res, err := l.svcCtx.StockModel.Insert(l.ctx, &model.Stock{
		GoodsId: in.GoodsId,
		Count:   in.Count,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &stock.CreateStockResp{Message: "创建成功", StockId: lastId}, nil
}
