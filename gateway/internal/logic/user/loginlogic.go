package user

import (
	"context"
	"gateway/service/user"
	"time"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	resp = &types.LoginResp{}
	u, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginReq{UserName: req.UserName, Password: req.Password})
	if err != nil {
		return nil, err
	}

	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	now := time.Now().Unix()
	accessToken, err := l.GenToken(now, l.svcCtx.Config.Auth.AccessSecret, nil, accessExpire)
	if err != nil {
		return nil, err
	}

	resp.User.Uid = u.User.ID
	resp.User.Username = u.User.UserName
	resp.AccessToken = accessToken
	resp.AccessExpire = now + accessExpire
	resp.RefreshAfter = now + accessExpire/2

	return resp, nil
}

func (l *LoginLogic) GenToken(iat int64, secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
