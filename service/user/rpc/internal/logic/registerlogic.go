package logic

import (
	"context"
	"fs-sys/common/cryptx"
	"fs-sys/service/user/rpc/internal/svc"
	"fs-sys/service/user/rpc/model"
	"fs-sys/service/user/rpc/user"
	"google.golang.org/grpc/status"
	"net/http"

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

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.UserInfoResponse, error) {
	// 判断 手机号是否注册过
	_, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err == nil {
		return nil, status.Error(100, "用户已经存在")
	}
	if err == model.ErrNotFound {
		newUserInfo := model.FsUser{
			Username: in.Username,
			Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
			Mobile:   in.Mobile,
		}
		res, err := l.svcCtx.UserModel.Insert(l.ctx, &newUserInfo)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
		userId, err := res.LastInsertId()
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
		return &user.UserInfoResponse{
			Id:       userId,
			Username: newUserInfo.Username,
			Mobile:   newUserInfo.Mobile,
		}, nil
	}
	return nil, status.Error(http.StatusInternalServerError, err.Error())
}
