package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TblUserFileModel = (*customTblUserFileModel)(nil)

type (
	// TblUserFileModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTblUserFileModel.
	TblUserFileModel interface {
		tblUserFileModel
	}

	customTblUserFileModel struct {
		*defaultTblUserFileModel
	}
)

// NewTblUserFileModel returns a model for the database table.
func NewTblUserFileModel(conn sqlx.SqlConn, c cache.CacheConf) TblUserFileModel {
	return &customTblUserFileModel{
		defaultTblUserFileModel: newTblUserFileModel(conn, c),
	}
}
