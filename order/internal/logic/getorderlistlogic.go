package logic

import (
	"context"

	"order/internal/svc"
	"order/service/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderListLogic {
	return &GetOrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderListLogic) GetOrderList(in *order.OrderListReq) (*order.OrderListResp, error) {
	// todo: add your logic here and delete this line

	return &order.OrderListResp{}, nil
}
