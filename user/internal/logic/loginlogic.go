package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"user/internal/model"
	"user/internal/svc"
	"user/service/user"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (out *user.UserResponse, err error) {
	out = &user.UserResponse{}
	out.Code = 200
	out.Message = "success"
	u, err := l.svcCtx.UserModel.FindOneByUsername(context.Background(), in.UserName)
	if err != nil {
		if err == model.ErrNotFound {
			out.Code = 400
			out.Message = "user not found"
		}
		out.Message = "unknown error"
		out.Code = 500
	}
	if CheckPassword(u, in.Password) == false {
		out.Code = 400
		out.Message = "username or password invalid"
	}
	out.User = model2Pb(u)

	return out, nil
}

func CheckPassword(u *model.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))
	return err == nil
}

func model2Pb(u *model.User) *user.UserModel {
	// todo: null value handle
	var updateAt, createAt int64
	if u.CreatedAt.Valid {
		createAt = u.CreatedAt.Time.Unix()
	}
	if u.UpdatedAt.Valid {
		updateAt = u.UpdatedAt.Time.Unix()
	}
	userModel := user.UserModel{
		ID:        u.Id,
		UserName:  u.Username,
		CreatedAt: createAt,
		UpdatedAt: updateAt,
	}
	return &userModel
}
