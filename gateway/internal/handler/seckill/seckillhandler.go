package seckill

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"gateway/internal/logic/seckill"
	"gateway/internal/svc"
	"gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SeckillHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SeckillReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error("handler error:", err)
			httpx.Error(w, err)
			return
		}

		l := seckill.NewSeckillLogic(r.Context(), svcCtx)
		resp, err := l.Seckill(&req)
		if err != nil {
			logx.Error("rpc call err: ", err)
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
