package logic

import (
	"context"
	"fmt"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"seckill/internal/svc"
	"seckill/service/order"
	"seckill/service/seckill"

	"github.com/zeromicro/go-zero/core/logx"
)

type Seckill2Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSeckill2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Seckill2Logic {
	return &Seckill2Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Seckill2 是二阶段消息模式实现秒杀，保证库存在redis与数据库之间的最终一致性、保证库存与订单之间的最终一致性。
// (1) 给order表新建一条订单记录、(2) 推送扣减消息到kafka异步更新stock表
// (1) 利用数据库的事务原子特性，要么完成要么不完成。(1) 原则上不允许失败，在使用dtm时，将重试这个步骤直至成功。
// (2) 无法使用数据库的回滚特性，但是由于(1)的存在，我们在消费队列消息时查询order表，根据order记录是否存在确定是否扣减stock, 这样就保证了数据最终一致性。
func (l *Seckill2Logic) Seckill2(in *seckill.SeckillReq) (*seckill.BaseSeckillResp, error) {
	orderServer, err := l.svcCtx.Config.OrderRpcConf.BuildTarget()
	if err != nil {
		logx.Error("build target fail: ", err)
		return nil, err
	}

	stockServer, err := l.svcCtx.Config.StockRpcConf.BuildTarget()
	if err != nil {
		logx.Error("build target fail: ", err)
		return nil, err
	}

	orderReq := order.OrderModel{
		Uid:        in.Uid,
		ActivityId: in.ActivityId,
		GoodsId:    in.GoodsId,
		StockId:    in.StockId,
		Count:      in.BuyCnt,
	}

	stockKey := fmt.Sprintf(l.svcCtx.Config.StockRedisKeyFormat, in.ActivityId, in.GoodsId, in.StockId)
	gid := stockKey + "-" + dtmgrpc.MustGenGid(l.svcCtx.Config.DtmServer.Addr)

	msg := dtmgrpc.NewMsgGrpc(l.svcCtx.Config.DtmServer.Addr, gid).
		Add(orderServer+"/order.Order/CreateOrderAndDeductAsync", &orderReq)

	msg.WaitResult = true
	msg.TimeoutToFail = 60

	err = msg.DoAndSubmit(stockServer+"/stock.stock/redisQueryPrepare", func(bb *dtmcli.BranchBarrier) error {
		err := bb.RedisCheckAdjustAmount(l.svcCtx.RawRedis, stockKey, -1, 86400)
		if err != nil {
			logx.Error(err, in)
			return status.Error(codes.Aborted, dtmcli.ResultFailure)
		} else {
			return nil
		}
	})

	if err != nil {
		logx.Error(err, in)
		return &seckill.BaseSeckillResp{Message: "fail"}, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	return &seckill.BaseSeckillResp{Message: "success"}, nil
}
