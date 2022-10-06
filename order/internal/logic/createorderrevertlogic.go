package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/client/dtmgrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"order/internal/model"
	"order/internal/svc"
	"order/service/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderRevertLogic {
	return &CreateOrderRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderRevertLogic) CreateOrderRevert(in *order.OrderModel) (*order.OrderModel, error) {
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		logx.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = barrier.CallWithDB(l.svcCtx.Mysql, func(tx *sql.Tx) error {
		err := l.svcCtx.OrderModel.UpdateTx(l.ctx, tx, &model.Order{
			Uid:        in.Uid,
			ActivityId: in.ActivityId,
			GoodsId:    in.GoodsId,
			StockId:    in.StockId,
			Status:     -1,
		})
		if err != nil {
			logx.Error(err)
			return status.Error(codes.Internal, err.Error())
		}

		return nil
	})
	if err != nil {
		logx.Errorf("create order revert err : %v with req: %+v\n", err, in)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &order.OrderModel{}, nil
}
