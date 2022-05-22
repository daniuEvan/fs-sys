package logic

import (
	"context"
	"fs-sys/service/user/rpc/model"
	"google.golang.org/grpc/status"

	"fs-sys/service/user/rpc/internal/svc"
	"fs-sys/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, err
	}
	return &user.UserInfoResponse{
		Id:         userInfo.Id,
		Username:   userInfo.Username,
		Mobile:     userInfo.Mobile,
		CreateTime: userInfo.CreateTime.Unix(),
		UpdateTime: userInfo.UpdateTime.Unix(),
	}, nil
}
