package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FsUserFileModel = (*customFsUserFileModel)(nil)

type (
	// FsUserFileModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFsUserFileModel.
	FsUserFileModel interface {
		fsUserFileModel
	}

	customFsUserFileModel struct {
		*defaultFsUserFileModel
	}
)

// NewFsUserFileModel returns a model for the database table.
func NewFsUserFileModel(conn sqlx.SqlConn, c cache.CacheConf) FsUserFileModel {
	return &customFsUserFileModel{
		defaultFsUserFileModel: newFsUserFileModel(conn, c),
	}
}
