package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"order/internal/model"

	"order/internal/svc"
	"order/service/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderLogic {
	return &UpdateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderLogic) UpdateOrder(in *order.OrderModel) (*order.BaseOrderResp, error) {
	err := l.svcCtx.OrderModel.Update(l.ctx, &model.Order{
		Id:         in.Id,
		Uid:        in.Uid,
		ActivityId: in.ActivityId,
		GoodsId:    in.GoodsId,
		StockId:    in.StockId,
		Count:      in.Count,
		Status:     in.Status,
		CreatedAt:  in.CreatedAt,
		UpdatedAt:  in.UpdatedAt,
		DeletedAt:  in.DeletedAt,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &order.BaseOrderResp{Message: "success"}, nil
}
