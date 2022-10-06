package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/client/dtmgrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"stock/internal/model"

	"stock/internal/svc"
	"stock/service/stock"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeductStockFastLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeductStockFastLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductStockFastLogic {
	return &DeductStockFastLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeductStockFastLogic) DeductStockFast(in *stock.DeductStockReq) (*stock.DeductStockResp, error) {
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		logx.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = barrier.CallWithDB(l.svcCtx.Mysql, func(tx *sql.Tx) error {
		m := &model.StockTask{
			StockId: in.StockId,
			Amount:  in.Count,
		}
		res, err := l.svcCtx.StockTaskModel.InsertWithTx(l.ctx, tx, m)
		if err == nil {
			affected, err := res.RowsAffected()
			lastInsertId, err2 := res.LastInsertId()
			if err == nil && err2 == nil {
				if affected == 1 {
					m.Id = lastInsertId
					err = SendDeductMessage(l.svcCtx.Producer, m, l.svcCtx.Config.StockDeductTopic)
					if err == nil {
						return nil
					} else {
						logx.Error(err)
					}
				} else {
					logx.Error("affected != 1")
				}
			} else {
				logx.Error(err, err2)
			}
		} else {
			logx.Error(err)
		}
		return status.Error(codes.Internal, err.Error()) // 不允许失败，标记为codes.Internal, 让dtm重试
	})
	if err != nil {
		logx.Error(err)
		return &stock.DeductStockResp{Message: err.Error()}, status.Error(codes.Internal, err.Error())
	} else {
		return &stock.DeductStockResp{Message: "success"}, nil
	}
}
