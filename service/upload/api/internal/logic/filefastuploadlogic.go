package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"fs-sys/common/fileHash"
	"fs-sys/service/upload/api/internal/model"
	"fs-sys/service/upload/rpc/upload"
	"io"
	"net/http"
	"time"

	"fs-sys/service/upload/api/internal/svc"
	"fs-sys/service/upload/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileFastUploadLogic struct {
	httpRequest *http.Request
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileFastUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *FileFastUploadLogic {
	return &FileFastUploadLogic{
		httpRequest: r,
		Logger:      logx.WithContext(ctx),
		ctx:         ctx,
		svcCtx:      svcCtx,
	}
}

//
// FileFastUpload
// @Description: 秒传逻辑
// @receiver l
// @param req:
// @return resp:
// @return err:
//
func (l *FileFastUploadLogic) FileFastUpload() (resp *types.UploadResponse, err error) {
	// 获取用户id
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		l.Logger.Error(fmt.Sprintf("获取用户id异常: %s", err.Error()))
		return nil, err
	}
	// 获取表单数据文件
	file, head, err := l.httpRequest.FormFile("file")
	if err != nil {
		l.Logger.Error("从表单获取文件失败")
		return nil, err
	}
	defer file.Close()
	// 把文件转化为[]byte 字节数组
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		l.Logger.Error("从表单获取文件失败")
		return nil, err
	}
	// 构建元数据信息
	fileMeta := model.FileMetadata{
		FileHash: fileHash.Sha1(buf.Bytes()),
		FileName: head.Filename,
		FileSize: int64(len(buf.Bytes())),
		UploadAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	// 1. 查询文件信息表是否已存在文件
	status, err := l.svcCtx.UploadRpc.FastUpload(l.ctx,
		&upload.FastUploadRequest{
			UserId: userId,
			FileMeta: &upload.FileMeta{
				FileHash: fileMeta.FileHash,
				FileName: fileMeta.FileName,
				FileSize: fileMeta.FileSize,
			},
		},
	)
	if err != nil {
		l.Logger.Errorf("秒传失败:%s", err.Error())
		return nil, err
	}
	if status.Status == 1 {
		return nil, errors.New("文件未找到,尝试直接上传")
	}
	return &types.UploadResponse{
		UserID:   userId,
		FileName: fileMeta.FileName,
		FileHash: fileMeta.FileHash,
		FileSize: fileMeta.FileSize,
	}, nil

}
