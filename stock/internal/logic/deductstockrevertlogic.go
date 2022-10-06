package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/client/dtmgrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"stock/internal/svc"
	"stock/service/stock"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeductStockRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeductStockRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductStockRevertLogic {
	return &DeductStockRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeductStockRevertLogic) DeductStockRevert(in *stock.DeductStockReq) (*stock.DeductStockResp, error) {
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		logx.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = barrier.CallWithDB(l.svcCtx.Mysql, func(tx *sql.Tx) error {
		query := "update stock set count=count+? where id=? and goods_id=?"
		res, err := tx.ExecContext(l.ctx, query, in.Count, in.StockId, in.GoodsId)
		if err == nil {
			affected, err := res.RowsAffected()
			if err == nil && affected == 1 {
				return nil
			} else {
				logx.Error(err)
			}
		} else {
			logx.Error(err)
		}
		return status.Error(codes.Internal, err.Error())
	})
	if err != nil {
		logx.Error(err)
		return &stock.DeductStockResp{Message: err.Error()}, status.Error(codes.Internal, err.Error())
	} else {
		return &stock.DeductStockResp{Message: "revert success"}, nil
	}
}
