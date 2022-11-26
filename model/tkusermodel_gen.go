// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tkUserFieldNames          = builder.RawFieldNames(&TkUser{})
	tkUserRows                = strings.Join(tkUserFieldNames, ",")
	tkUserRowsExpectAutoSet   = strings.Join(stringx.Remove(tkUserFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	tkUserRowsWithPlaceHolder = strings.Join(stringx.Remove(tkUserFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheTkUserIdPrefix = "cache:tkUser:id:"
)

type (
	tkUserModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *TkUser) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TkUser, error)
		Update(ctx context.Context, session sqlx.Session, newData *TkUser) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultTkUserModel struct {
		sqlc.CachedConn
		table string
	}

	TkUser struct {
		Id        int64        `db:"id"`
		Account   string       `db:"account"`
		Password  string       `db:"password"`
		CreatedAt time.Time    `db:"created_at"`
		UpdatedAt sql.NullTime `db:"updated_at"`
	}
)

func newTkUserModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultTkUserModel {
	return &defaultTkUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`tk_user`",
	}
}

func (m *defaultTkUserModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	tkUserIdKey := fmt.Sprintf("%s%v", cacheTkUserIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, tkUserIdKey)
	return err
}

func (m *defaultTkUserModel) FindOne(ctx context.Context, id int64) (*TkUser, error) {
	tkUserIdKey := fmt.Sprintf("%s%v", cacheTkUserIdPrefix, id)
	var resp TkUser
	err := m.QueryRowCtx(ctx, &resp, tkUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tkUserRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTkUserModel) Insert(ctx context.Context, session sqlx.Session, data *TkUser) (sql.Result, error) {
	tkUserIdKey := fmt.Sprintf("%s%v", cacheTkUserIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, tkUserRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Account, data.Password, data.CreatedAt, data.UpdatedAt)
		}
		return conn.ExecCtx(ctx, query, data.Account, data.Password, data.CreatedAt, data.UpdatedAt)
	}, tkUserIdKey)
	return ret, err
}

func (m *defaultTkUserModel) Update(ctx context.Context, session sqlx.Session, data *TkUser) error {
	tkUserIdKey := fmt.Sprintf("%s%v", cacheTkUserIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tkUserRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Account, data.Password, data.CreatedAt, data.UpdatedAt, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.Account, data.Password, data.CreatedAt, data.UpdatedAt, data.Id)
	}, tkUserIdKey)
	return err
}

func (m *defaultTkUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheTkUserIdPrefix, primary)
}

func (m *defaultTkUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tkUserRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTkUserModel) tableName() string {
	return m.table
}
