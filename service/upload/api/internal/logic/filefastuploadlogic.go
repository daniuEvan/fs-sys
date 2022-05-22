package logic

import (
	"context"

	"fs-sys/service/upload/api/internal/svc"
	"fs-sys/service/upload/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileFastUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileFastUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileFastUploadLogic {
	return &FileFastUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileFastUploadLogic) FileFastUpload(req *types.FastUploadReq) (resp *types.UploadResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
