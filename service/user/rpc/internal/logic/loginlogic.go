package logic

import (
	"context"
	"fs-sys/common/cryptx"
	"fs-sys/service/user/rpc/model"
	"google.golang.org/grpc/status"

	"fs-sys/service/user/rpc/internal/svc"
	"fs-sys/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.UserInfoResponse, error) {
	res, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	// 判断密码是否正确
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if password != res.Password {
		return nil, status.Error(100, "密码错误")
	}
	return &user.UserInfoResponse{
		Id:         res.Id,
		Username:   res.Username,
		Mobile:     res.Mobile,
		CreateTime: res.CreateTime.Unix(),
		UpdateTime: res.UpdateTime.Unix(),
	}, nil
}
