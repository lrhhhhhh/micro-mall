package user

import (
	"gateway/internal/types"
	"net/http"

	"gateway/internal/logic/user"
	"gateway/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewGetUserLogic(r.Context(), svcCtx)
		resp, err := l.GetUser(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
