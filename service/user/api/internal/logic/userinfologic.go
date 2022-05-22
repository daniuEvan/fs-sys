package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fs-sys/service/user/api/internal/svc"
	"fs-sys/service/user/api/internal/types"
	"fs-sys/service/user/rpc/user"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	userId, ok := l.ctx.Value("userId").(json.Number)
	if !ok {
		return nil, errors.New("token 异常")
	}
	userIdInt64, err := userId.Int64()
	userInfo, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: userIdInt64,
	})
	return &types.UserInfoResponse{
		ID:         userInfo.Id,
		Username:   userInfo.Username,
		Mobile:     userInfo.Mobile,
		CreateTime: userInfo.CreateTime,
		UpdateTime: userInfo.UpdateTime,
	}, err
}
