package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"stock/internal/svc"
	"stock/service/stock"
	"time"
)

type DeductStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeductStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductStockLogic {
	return &DeductStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeductStockLogic) DeductStock(in *stock.DeductStockReq) (*stock.DeductStockResp, error) {
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		logx.Error(err)
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure) // 立即失败
	}

	dl, _ := l.ctx.Deadline()
	err = barrier.CallWithDB(l.svcCtx.Mysql, func(tx *sql.Tx) error {
		query := "update stock set count=count-? where id=? and goods_id=? and count >= ?"
		session := sqlx.NewSessionFromTx(tx)
		res, err := session.ExecCtx(l.ctx, query, in.Count, in.StockId, in.GoodsId, in.Count)
		//res, err := tx.ExecContext(l.ctx, query, in.Count, in.StockId, in.GoodsId, in.Count)
		if err == nil {
			affected, err := res.RowsAffected()
			if err == nil {
				if affected == 1 {
					return nil
				} else {
					logx.Error("must one row affected")
				}
			} else {
				logx.Error(err)
			}
		} else {
			logx.Error(err)
		}
		return status.Error(codes.Aborted, dtmcli.ResultFailure) // 立即失败，不重试
	})
	if err != nil {
		logx.Error(err, "deadline: ", time.Until(dl).String())
		return &stock.DeductStockResp{Message: err.Error()}, status.Error(codes.Aborted, dtmcli.ResultFailure)
	} else {
		return &stock.DeductStockResp{Message: "success"}, nil
	}
}
