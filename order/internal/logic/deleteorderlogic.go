package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"order/internal/svc"
	"order/service/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderLogic {
	return &DeleteOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteOrderLogic) DeleteOrder(in *order.OrderModel) (*order.BaseOrderResp, error) {
	err := l.svcCtx.OrderModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &order.BaseOrderResp{}, nil
}
