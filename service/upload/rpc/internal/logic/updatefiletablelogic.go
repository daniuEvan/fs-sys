package logic

import (
	"context"
	"fs-sys/service/upload/rpc/internal/svc"
	"fs-sys/service/upload/rpc/model"
	"fs-sys/service/upload/rpc/upload"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFileTableLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFileTableLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFileTableLogic {
	return &UpdateFileTableLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateFileTable 更新文件表 todo 与更新用户表的事务问题
func (l *UpdateFileTableLogic) UpdateFileTable(in *upload.FileMeta) (*upload.Empty, error) {
	filesInfo := model.FsFiles{
		FileHash: in.FileHash,
		FileName: in.FileName,
		FileSize: in.FileSize,
		FileAddr: in.Location,
		Status:   0,
	}

	_, err := l.svcCtx.FilesTableModel.FindOneByFileHash(l.ctx, in.FileHash)
	if err == model.ErrNotFound {
		_, err = l.svcCtx.FilesTableModel.Insert(l.ctx, &filesInfo)
		if err != nil {
			l.Logger.Errorf("文件信息表插入失败: ", err)
			return &upload.Empty{}, err
		}
	} else if err != nil {
		l.Logger.Errorf("文件信息表查询失败: ", err)
		return &upload.Empty{}, err
	}
	err = l.svcCtx.FilesTableModel.Update(l.ctx, &filesInfo)
	if err != nil {
		l.Logger.Errorf("文件信息表更新失败: ", err)
		return &upload.Empty{}, err
	}
	return &upload.Empty{}, nil
}
