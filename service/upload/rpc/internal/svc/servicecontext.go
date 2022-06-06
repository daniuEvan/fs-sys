package svc

import (
	"fs-sys/service/upload/rpc/internal/config"
	"fs-sys/service/upload/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config          config.Config
	FilesTableModel model.FsFilesModel
	FileUserModel   model.FsUserFileModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:          c,
		FilesTableModel: model.NewFsFilesModel(conn, c.CacheRedis),
		FileUserModel:   model.NewFsUserFileModel(conn, c.CacheRedis),
	}
}
