package model

import (
	"context"
	"github.com/Masterminds/squirrel"
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
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
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

// export logic
func (c *customTkUserModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(tkUserRows).From(c.table)
}

// export logic
func (c *customTkUserModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(c.table)
}

// export logic
func (c *customTkUserModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(c.table)
}
