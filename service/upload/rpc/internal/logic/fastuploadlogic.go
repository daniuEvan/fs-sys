package logic

import (
	"context"
	"fs-sys/service/upload/rpc/internal/svc"
	"fs-sys/service/upload/rpc/model"
	"fs-sys/service/upload/rpc/upload"
	"github.com/zeromicro/go-zero/core/logx"
)

type FastUploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFastUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FastUploadLogic {
	return &FastUploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FastUpload 秒传
func (l *FastUploadLogic) FastUpload(in *upload.FastUploadRequest) (*upload.FastUploadResponse, error) {
	fileHash := in.FileMeta.FileHash
	// 查询文件列表
	_, err := l.svcCtx.FilesTableModel.FindOneByFileHash(l.ctx, fileHash)
	if err == model.ErrNotFound {
		l.Logger.Info("未查询到文件")
		return &upload.FastUploadResponse{Status: 1}, nil
	} else if err != nil {
		l.Logger.Errorf("查询files table失败:%s", err.Error())
		return &upload.FastUploadResponse{Status: 1}, err
	}
	// 查询用户文件表
	userId := in.UserId
	fileName := in.FileMeta.FileName
	data := model.FsUserFile{
		UserId:   userId,
		FileHash: fileHash,
		FileSize: in.FileMeta.FileSize,
		FileName: fileName,
		Status:   0,
	}
	fileUserTableRes, err := l.svcCtx.FileUserModel.FindOneByUserIdFileHashFileName(l.ctx, userId, fileHash, fileName)
	if err == model.ErrNotFound {
		_, err = l.svcCtx.FileUserModel.Insert(l.ctx, &data)
		if err != nil {
			l.Logger.Errorf("插入 user file table 失败:%s", err.Error())
			return &upload.FastUploadResponse{Status: 1}, err
		}
		return &upload.FastUploadResponse{Status: 0}, nil
	} else if err != nil {
		l.Logger.Errorf("查询 user file table失败:%s", err.Error())
		return &upload.FastUploadResponse{Status: 1}, err
	}
	data.Id = fileUserTableRes.Id
	err = l.svcCtx.FileUserModel.Update(l.ctx, &data)
	if err != nil {
		l.Logger.Errorf("更新 user file table 失败:%s", err.Error())
		return &upload.FastUploadResponse{Status: 1}, err
	}
	return &upload.FastUploadResponse{Status: 0}, nil
}
