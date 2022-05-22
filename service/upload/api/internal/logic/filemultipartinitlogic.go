package logic

import (
	"context"

	"fs-sys/service/upload/api/internal/svc"
	"fs-sys/service/upload/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileMultipartInitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileMultipartInitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileMultipartInitLogic {
	return &FileMultipartInitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileMultipartInitLogic) FileMultipartInit(req *types.MultipartUploadInitReq) (resp *types.MultipartUploadInitRes, err error) {
	// todo: add your logic here and delete this line

	return
}
