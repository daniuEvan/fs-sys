package logic

import (
	"context"
	"fs-sys/common/jwtx"
	"fs-sys/service/user/rpc/user"
	"time"

	"fs-sys/service/user/api/internal/svc"
	"fs-sys/service/user/api/internal/types"

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

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResonse, err error) {
	userObj, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginRequest{
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	nowTime := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	accessToken, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, nowTime, accessExpire, userObj.Id)
	if err != nil {
		return nil, err
	}

	return &types.LoginResonse{
		AccessToken:  accessToken,
		AccessExpire: nowTime + accessExpire,
	}, err

	return
}
