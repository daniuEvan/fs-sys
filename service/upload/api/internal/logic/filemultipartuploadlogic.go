package logic

import (
	"context"

	"fs-sys/service/upload/api/internal/svc"
	"fs-sys/service/upload/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileMultipartUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileMultipartUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileMultipartUploadLogic {
	return &FileMultipartUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileMultipartUploadLogic) FileMultipartUpload(req *types.MultipartUploadReq) error {
	// todo: add your logic here and delete this line

	return nil
}
