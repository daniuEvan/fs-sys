package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"fs-sys/common/fileHash"
	"fs-sys/common/utils"
	"fs-sys/service/upload/api/internal/model"
	"fs-sys/store/minioStore"
	"github.com/minio/minio-go/v7"
	"path"
	"strconv"

	//"fs-sys/store/minioStore"
	//"github.com/minio/minio-go/v7"
	"io"
	"net/http"
	"os"
	"time"

	"fs-sys/service/upload/api/internal/svc"
	"fs-sys/service/upload/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	httpRequest *http.Request
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *FileUploadLogic {
	return &FileUploadLogic{
		httpRequest: r,
		Logger:      logx.WithContext(ctx),
		ctx:         ctx,
		svcCtx:      svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload() (resp *types.UploadResponse, err error) {

	// 1. 从表单获取文件
	file, head, err := l.httpRequest.FormFile("file")
	if err != nil {
		l.Logger.Error("从表单获取文件失败")
		return nil, err
	}
	defer file.Close()
	// 2. 把文件内容转为[]byte
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		l.Logger.Error("从表单获取文件失败")
		return nil, err
	}
	// 3. 构建文件元信息
	fileMeta := model.FileMetadata{
		FileHash: fileHash.Sha1(buf.Bytes()),
		FileName: head.Filename,
		FileSize: int64(len(buf.Bytes())),
		UploadAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	// 4. 将文件写入临时存储位置

	fileLocation := path.Join(
		l.svcCtx.Config.Uploader.FileTempLocation,
		fmt.Sprintf("%s-%s", fileMeta.FileHash, fileMeta.FileName),
	)
	fileMeta.Location = fileLocation
	newFile, err := os.Create(fileMeta.Location)
	if err != nil {
		l.Logger.Error(fmt.Sprintf("文件创建失败: %s", err.Error()))
		return nil, err
	}
	defer newFile.Close()
	nByte, err := newFile.Write(buf.Bytes())
	if int64(nByte) != fileMeta.FileSize || err != nil {
		l.Logger.Error(fmt.Sprintf("文件写入失败: %s", err.Error()))
		return nil, err
	}
	newFile.Seek(0, 0)
	// 判断文件content-type
	fileContentType, err := utils.GetFileContentType(newFile)
	if err != nil {
		l.Logger.Error(fmt.Sprintf("文件类型判断失败: %s", err.Error()))
		return nil, err
	}
	fmt.Println(fileContentType)
	// 5. 同步或异步将文件转移到minio OSS
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
	bucketName := l.svcCtx.Config.Uploader.BucketName
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
	// 判断写入oss是同步(直接上传)还是异步(写入异步队列中)
	if !l.svcCtx.Config.Uploader.AsyncUpload {
		info, err := minioClient.FPutObject(
			l.ctx,
			bucketName,
			fileMeta.FileName,
			fileMeta.Location,
			minio.PutObjectOptions{ContentType: fileContentType},
		)
		if err != nil {
			l.Logger.Error(fmt.Sprintf("minio 文件创建失败: %s", err.Error()))
			return nil, err
		}
		fmt.Println(info)
	}
	// 6. 更新文件表记录
	// 7. 更新用户文件表

	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		l.Logger.Error(fmt.Sprintf("登录异常: %s", err.Error()))
		return nil, err
	}
	//userId = json.Number(int64(userId))
	return &types.UploadResponse{
		UserID:   userId,
		FileName: fileMeta.FileName,
		FileHash: fileMeta.FileHash,
		FileSize: strconv.Itoa(int(fileMeta.FileSize)),
	}, nil
}
