package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"order/internal/svc"
	"order/service/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderLogic) GetOrder(in *order.OrderModel) (*order.OrderModel, error) {
	res, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &order.OrderModel{
		Id:         res.Id,
		Uid:        res.Uid,
		ActivityId: res.ActivityId,
		GoodsId:    res.GoodsId,
		StockId:    res.StockId,
		Count:      res.Count,
		Status:     res.Status,
		CreatedAt:  res.CreatedAt,
		UpdatedAt:  res.UpdatedAt,
		DeletedAt:  res.DeletedAt,
	}, nil
}
