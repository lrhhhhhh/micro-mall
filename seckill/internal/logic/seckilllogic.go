package logic

import (
	"context"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"math/rand"
	"seckill/service/order"
	"seckill/service/stock"
	"time"

	"seckill/internal/svc"
	"seckill/service/seckill"

	"github.com/zeromicro/go-zero/core/logx"
)

type SeckillLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSeckillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillLogic {
	return &SeckillLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Seckill 是TCC模式实现的，保证数据库中订单和库存的数据最终一致性（包含回滚取消的订单），不保证redis和库存的数据一致性。
func (l *SeckillLogic) Seckill(in *seckill.SeckillReq) (*seckill.BaseSeckillResp, error) {
	activityName := fmt.Sprintf("SeckillActivity:%d", in.ActivityId)
	historyName := fmt.Sprintf("SeckillHistory:%d", in.ActivityId)
	uid := in.Uid
	buyCnt := in.BuyCnt
	now := time.Now().Unix()
	probability := rand.Float64()

	code, err := l.svcCtx.Redis.EvalShaCtx(
		l.ctx,
		l.svcCtx.Config.StockDeductSha1,
		[]string{activityName, historyName},
		[]interface{}{uid, buyCnt, now, probability},
	)
	if err != nil {
		logx.Error(err, l.svcCtx.Config.StockDeductSha1)
		return nil, status.Error(codes.Internal, err.Error())
	}

	out := &seckill.BaseSeckillResp{}
	out.Code = code.(int64)
	out.Message = GetMessage(int(out.Code))
	logx.Infof("redis stock deduct resp: (code=%d, message=%s)", out.Code, out.Message)

	if int(code.(int64)) != SeckillSuccess {
		return out, nil
	}

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

	stockReq := stock.DeductStockReq{
		StockId: in.StockId,
		GoodsId: in.GoodsId,
		Count:   in.BuyCnt,
	}

	var d time.Duration = 0
	parentDeadline, _ := l.ctx.Deadline()
	deadline := time.Until(parentDeadline) - time.Millisecond*100
	if deadline > d {
		d = deadline
	}
	ctx, cancel := context.WithTimeout(l.ctx, d)
	custom := func(tcc *dtmgrpc.TccGrpc) {
		tcc.Context = ctx
	}
	defer cancel()

	gid := dtmgrpc.MustGenGid(l.svcCtx.Config.DtmServer.Addr)
	err = dtmgrpc.TccGlobalTransaction2(l.svcCtx.Config.DtmServer.Addr, gid, custom, func(tcc *dtmgrpc.TccGrpc) error {
		err := tcc.CallBranch(
			&stockReq,
			stockServer+"/stock.stock/deductStock",
			"",
			stockServer+"/stock.stock/deductStockRevert",
			&emptypb.Empty{},
		)
		if err != nil {
			return err
		}

		err = tcc.CallBranch(
			&orderReq,
			orderServer+"/order.Order/CreateOrder",
			"",
			orderServer+"/order.Order/CreateOrderRevert",
			&emptypb.Empty{},
		)

		return err
	})

	if err != nil {
		logx.Error(err, in)
		return &seckill.BaseSeckillResp{Message: "fail"}, err
	}
	return &seckill.BaseSeckillResp{Message: "success"}, nil
}
