package svc

import (
	"fs-sys/service/upload/api/internal/config"
	"fs-sys/service/upload/rpc/upload"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	UploadRpc upload.Upload
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UploadRpc: upload.NewUpload(zrpc.MustNewClient(c.UploadRpc)),
	}
}
