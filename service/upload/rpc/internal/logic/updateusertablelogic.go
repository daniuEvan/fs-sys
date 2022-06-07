package logic

import (
	"context"
	"fs-sys/service/upload/rpc/internal/svc"
	"fs-sys/service/upload/rpc/model"
	"fs-sys/service/upload/rpc/upload"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserTableLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserTableLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserTableLogic {
	return &UpdateUserTableLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateUserTable 更新用户表
func (l *UpdateUserTableLogic) UpdateUserTable(in *upload.UserTableUpdateRequest) (*upload.Empty, error) {
	data := model.FsUserFile{
		UserId:   in.UserId,
		FileHash: in.FileMeta.FileHash,
		FileSize: in.FileMeta.FileSize,
		FileName: in.FileMeta.FileName,
		Status:   0,
	}
	fileUserTableRes, err := l.svcCtx.FileUserModel.FindOneByUserIdFileHashFileName(l.ctx, in.UserId, in.FileMeta.FileHash, in.FileMeta.FileName)
	if err == model.ErrNotFound {
		_, err = l.svcCtx.FileUserModel.Insert(l.ctx, &data)
		if err != nil {
			l.Logger.Errorf("用户文件表插入失败: ", err)
			return &upload.Empty{}, err
		}
	} else if err != nil {
		l.Logger.Errorf("用户文件表查询失败: ", err)
		return &upload.Empty{}, err
	}
	data.Id = fileUserTableRes.Id
	err = l.svcCtx.FileUserModel.Update(l.ctx, &data)
	if err != nil {
		l.Logger.Errorf("更新用户文件表失败: ", err)
		return &upload.Empty{}, err
	}
	return &upload.Empty{}, nil
}
