package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TkUserModel = (*customTkUserModel)(nil)

type (
	// TkUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTkUserModel.
	TkUserModel interface {
		tkUserModel
	}

	customTkUserModel struct {
		*defaultTkUserModel
	}
)

// NewTkUserModel returns a model for the database table.
func NewTkUserModel(conn sqlx.SqlConn, c cache.CacheConf) TkUserModel {
	return &customTkUserModel{
		defaultTkUserModel: newTkUserModel(conn, c),
	}
}
