package logic

import (
	"context"
	"fs-sys/service/user/rpc/user"

	"fs-sys/service/user/api/internal/svc"
	"fs-sys/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.UserInfoResponse, err error) {
	res, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
		Mobile:   req.Mobile,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserInfoResponse{
		ID:         res.Id,
		Username:   res.Username,
		Mobile:     res.Mobile,
		CreateTime: res.CreateTime,
		UpdateTime: res.UpdateTime,
	}, nil

	return
}
