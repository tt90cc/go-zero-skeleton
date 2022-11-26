package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TkUserModel = (*customTkUserModel)(nil)

type (
	// TkUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTkUserModel.
	TkUserModel interface {
		tkUserModel
		Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
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

func (c *customTkUserModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return c.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}
