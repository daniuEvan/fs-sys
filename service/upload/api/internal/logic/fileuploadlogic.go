package logic

import (
	"bytes"
	"context"
	"fs-sys/common/fileHash"
	"fs-sys/service/upload/api/internal/model"
	"io"
	"net/http"
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
	fileMeta.Location =
	// 5. 同步或异步将文件转移到minio OSS
	// 6. 更新文件表记录
	// 7. 更新用户文件表
	return
}
