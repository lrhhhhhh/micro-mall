package logic

import (
	"context"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
	"user/internal/model"

	"user/internal/svc"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (out *user.UserResponse, err error) {
	out = &user.UserResponse{}
	if in.Password != in.PasswordConfirm {
		out.Code = http.StatusBadRequest
		out.Message = "两次密码输入不一致"
		return out, nil
	}
	u, err := l.svcCtx.UserModel.FindOneByUsername(context.Background(), in.UserName)
	if u != nil {
		out.Code = http.StatusBadRequest
		out.Message = "用户名已存在"
		return out, nil
	}

	newUser := model.User{
		Username:  in.UserName,
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}
	err = SetPassword(&newUser, in.Password)
	if err != nil {
		out.Code = http.StatusBadRequest
		out.Message = err.Error()
		return out, nil
	}

	res, err := l.svcCtx.UserModel.Insert(context.Background(), &newUser)
	if err != nil {
		out.Code = http.StatusBadRequest
		out.Message = err.Error()
		return out, nil
	}

	out.Code = 200
	out.Message = "success"
	id, _ := res.LastInsertId()
	newUser.Id = uint64(id) // todo: int to uint convert error
	out.User = model2Pb(&newUser)

	return out, nil
}

const (
	PassWordCost = 12
)

func SetPassword(u *model.User, password string) error {
	var bytes, err = bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	u.PasswordDigest = string(bytes)
	return nil
}
