package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FsFilesModel = (*customFsFilesModel)(nil)

type (
	// FsFilesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFsFilesModel.
	FsFilesModel interface {
		fsFilesModel
	}

	customFsFilesModel struct {
		*defaultFsFilesModel
	}
)

// NewFsFilesModel returns a model for the database table.
func NewFsFilesModel(conn sqlx.SqlConn, c cache.CacheConf) FsFilesModel {
	return &customFsFilesModel{
		defaultFsFilesModel: newFsFilesModel(conn, c),
	}
}
