package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/client/dtmgrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"order/internal/model"
	"time"

	"order/internal/svc"
	"order/service/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderLogic) CreateOrder(in *order.OrderModel) (*order.OrderModel, error) {
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		logx.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	var lastId int64
	now := time.Now().Unix()
	err = barrier.CallWithDB(l.svcCtx.Mysql, func(tx *sql.Tx) error {
		res, err := l.svcCtx.OrderModel.InsertTx(l.ctx, tx, &model.Order{
			Id:         0,
			Uid:        in.Uid,
			ActivityId: in.ActivityId,
			GoodsId:    in.GoodsId,
			StockId:    in.StockId,
			Count:      in.Count,
			Status:     in.Status,
			CreatedAt:  now,
			UpdatedAt:  now,
			DeletedAt:  0,
		})
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		lastId, err = res.LastInsertId()
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		logx.Infof("send <order:%d> to delay queue", lastId)
		err = CancelOrderAfter15m(l.svcCtx.Producer, int(lastId), 5, "order-cancel", "")
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		logx.Error(err, in)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &order.OrderModel{
		Id:         lastId,
		Uid:        in.Uid,
		ActivityId: in.ActivityId,
		GoodsId:    in.GoodsId,
		StockId:    in.StockId,
		Count:      in.Count,
		Status:     in.Status,
		CreatedAt:  now,
		UpdatedAt:  now,
		DeletedAt:  0,
	}, nil
}
