package logic

import (
	"context"

	"fs-sys/service/upload/api/internal/svc"
	"fs-sys/service/upload/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileMultipartMergeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileMultipartMergeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileMultipartMergeLogic {
	return &FileMultipartMergeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileMultipartMergeLogic) FileMultipartMerge(req *types.MultipartUploadMergeReq) error {
	// todo: add your logic here and delete this line

	return nil
}
