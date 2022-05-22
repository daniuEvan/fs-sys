package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FsUserModel = (*customFsUserModel)(nil)

type (
	// FsUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFsUserModel.
	FsUserModel interface {
		fsUserModel
	}

	customFsUserModel struct {
		*defaultFsUserModel
	}
)

// NewFsUserModel returns a model for the database table.
func NewFsUserModel(conn sqlx.SqlConn, c cache.CacheConf) FsUserModel {
	return &customFsUserModel{
		defaultFsUserModel: newFsUserModel(conn, c),
	}
}
