package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"fs-sys/common/fileHash"
	"fs-sys/common/utils"
	"fs-sys/service/upload/api/internal/model"
	"fs-sys/service/upload/rpc/upload"
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
	fileMeta.FileContentType = fileContentType
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		l.Logger.Error(fmt.Sprintf("获取用户id异常: %s", err.Error()))
		return nil, err
	}
	// 5. 同步或异步将文件转移到minio OSS
	// 判断写入oss是同步(直接上传)还是异步(写入异步队列中)
	if !l.svcCtx.Config.Uploader.AsyncUpload {
		info, err := l.svcCtx.UploadRpc.FileUpload(l.ctx, &upload.FileUploadRequest{
			FileMeta: &upload.FileMeta{
				FileHash:        fileMeta.FileHash,
				FileName:        path.Join(time.Now().Format("2006-01-02"), fmt.Sprintf("%d", userId), fileMeta.FileName),
				FileContentType: fileMeta.FileContentType,
				FileSize:        fileMeta.FileSize,
				Location:        fileMeta.Location,
				UploadAt:        fileMeta.UploadAt,
			},
			FileOSSMeta: &upload.FileOSSMeta{
				BucketName: l.svcCtx.Config.Uploader.BucketName,
				OssPath:    "", // todo oss 具体路径
			},
		})
		if err != nil {
			l.Logger.Error(fmt.Sprintf("文件上传失败: %s", err.Error()))
			return nil, err
		}
		//l.Logger.Info(info)
		fileBucketName := info.Bucket
		fileKey := info.Key
		// 6. 更新文件表记录
		_, err = l.svcCtx.UploadRpc.UpdateFileTable(l.ctx, &upload.FileMeta{
			FileHash: fileMeta.FileHash,
			FileName: fileMeta.FileName,
			FileSize: fileMeta.FileSize,
			Location: path.Join(fileBucketName, fileKey),
		})
		if err != nil {
			l.Logger.Error(fmt.Sprintf("文件表更新失败: %s", err.Error()))
			return nil, err
		}
		// 7. 更新用户文件表
		_, err = l.svcCtx.UploadRpc.UpdateUserTable(l.ctx, &upload.UserTableUpdateRequest{
			UserId: userId,
			FileMeta: &upload.FileMeta{
				FileHash: fileMeta.FileHash,
				FileName: fileMeta.FileName,
				FileSize: fileMeta.FileSize,
				Location: path.Join(fileBucketName, fileKey),
			},
			FileOSSMeta: &upload.FileOSSMeta{
				BucketName: fileBucketName,
				OssPath:    path.Join(fileBucketName, fileKey),
			},
		})
		if err != nil {
			l.Logger.Error(fmt.Sprintf("文件表更新失败: %s", err.Error()))
			return nil, err
		}

	} else { // 异步上传

		// 6. 更新文件表记录

		// 7. 更新用户文件表
	}

	//userId = json.Number(int64(userId))
	return &types.UploadResponse{
		UserID:   userId,
		FileName: fileMeta.FileName,
		FileHash: fileMeta.FileHash,
		FileSize: strconv.Itoa(int(fileMeta.FileSize)),
	}, nil
}
