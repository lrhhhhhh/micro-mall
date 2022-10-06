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

type CreateOrderAndDeductAsyncLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderAndDeductAsyncLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderAndDeductAsyncLogic {
	return &CreateOrderAndDeductAsyncLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateOrderAndDeductAsync 创建订单，发送取消订单消息到延迟队列，发送扣减库存消息到队列
// 只要redis扣减成功，那么创建订单一定要成功，所以以下实现都要重试到成功为止
func (l *CreateOrderAndDeductAsyncLogic) CreateOrderAndDeductAsync(in *order.OrderModel) (*order.OrderModel, error) {
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

		err = CancelOrderAfter15m(l.svcCtx.Producer, int(lastId), 5, "order-cancel", "")
		if err != nil {
			return err
		}

		err = SendDeductMessage(l.svcCtx.Producer, int(lastId), "stock-deduct")
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
