package logic

import (
	"context"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"stock/internal/svc"
	"stock/service/stock"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedisQueryPrepareLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRedisQueryPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedisQueryPrepareLogic {
	return &RedisQueryPrepareLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RedisQueryPrepareLogic) RedisQueryPrepare(in *stock.DeductStockReq) (*stock.DeductStockResp, error) {
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		logx.Error(err)
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}

	err = barrier.RedisQueryPrepared(l.svcCtx.Redis, 7*24*60*60)
	if err != nil {
		logx.Error("prepare err: ", err)
		return nil, err
	} else {
		return &stock.DeductStockResp{Message: "prepare done"}, nil
	}
}
