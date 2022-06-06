package logic

import (
	"context"
	"fmt"
	"fs-sys/store/minioStore"
	"github.com/minio/minio-go/v7"

	"fs-sys/service/upload/rpc/internal/svc"
	"fs-sys/service/upload/rpc/upload"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FileUpload 文件上传
func (l *FileUploadLogic) FileUpload(in *upload.FileUploadRequest) (*upload.FileUploadResponse, error) {
	// 获取minio客户端
	minioClient, err := minioStore.NewMinioClient(
		l.svcCtx.Config.Minio.Endpoint,
		l.svcCtx.Config.Minio.AccessKey,
		l.svcCtx.Config.Minio.SecretKey,
		l.svcCtx.Config.Minio.UseSSL,
	)
	if err != nil {
		l.Logger.Error(fmt.Sprintf("连接minio server 失败: %s", err.Error()))
		return nil, err
	}
	// 判断是否存在Bucket, 不存在创建  todo 完善bucket 命名逻辑
	bucketName := in.FileOSSMeta.BucketName
	if len(bucketName) < 1 {
		bucketName = l.svcCtx.Config.Minio.BucketName

	}
	found, err := minioClient.BucketExists(l.ctx, bucketName)
	if err != nil {
		l.Logger.Error(fmt.Sprintf("minio bucket 查询失败: %s", err.Error()))
		return nil, err
	}
	if !found {
		err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: "ch-cd", ObjectLocking: false})
		if err != nil {
			l.Logger.Error(fmt.Sprintf("minio bucket 创建失败: %s", err.Error()))
			return nil, err
		}
	}
	info, err := minioClient.FPutObject(
		l.ctx,
		bucketName,
		in.FileMeta.FileName,
		in.FileMeta.Location,
		minio.PutObjectOptions{ContentType: in.FileMeta.FileContentType},
	)
	if err != nil {
		l.Logger.Error(fmt.Sprintf("minio 文件创建失败: %s", err.Error()))
		return nil, err
	}

	return &upload.FileUploadResponse{
		Bucket:           info.Bucket,
		Key:              info.Key,
		ETag:             info.ETag,
		Size:             info.Size,
		LastModified:     info.LastModified.String(),
		Location:         info.Location,
		VersionID:        info.VersionID,
		Expiration:       info.Expiration.String(),
		ExpirationRuleID: info.ExpirationRuleID,
	}, nil
}
