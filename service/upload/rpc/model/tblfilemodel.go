package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TblFileModel = (*customTblFileModel)(nil)

type (
	// TblFileModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTblFileModel.
	TblFileModel interface {
		tblFileModel
	}

	customTblFileModel struct {
		*defaultTblFileModel
	}
)

// NewTblFileModel returns a model for the database table.
func NewTblFileModel(conn sqlx.SqlConn, c cache.CacheConf) TblFileModel {
	return &customTblFileModel{
		defaultTblFileModel: newTblFileModel(conn, c),
	}
}
