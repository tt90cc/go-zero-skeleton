package model

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
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
		InsertBuilder() squirrel.InsertBuilder
		UpdateBuilder() squirrel.UpdateBuilder
		DeleteBuilder() squirrel.DeleteBuilder
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*TkUser, error)
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*TkUser, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*TkUser, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*TkUser, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*TkUser, error)
		InsertBatch(ctx context.Context, session sqlx.Session, insertBuilder squirrel.InsertBuilder) (sql.Result, error)
		UpdateBatch(ctx context.Context, session sqlx.Session, updateBuilder squirrel.UpdateBuilder) (sql.Result, error)
		DeleteBatch(ctx context.Context, session sqlx.Session, deleteBuilder squirrel.DeleteBuilder) (sql.Result, error)
	}

	customTkUserModel struct {
		*defaultTkUserModel
	}
)

// NewTkUserModel returns a model for the database table.
func NewTkUserModel(conn sqlx.SqlConn) TkUserModel {
	return &customTkUserModel{
		defaultTkUserModel: newTkUserModel(conn),
	}
}

func (c *customTkUserModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return c.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
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

func (c *customTkUserModel) InsertBuilder() squirrel.InsertBuilder {
	return squirrel.Insert(c.table)
}

func (c *customTkUserModel) UpdateBuilder() squirrel.UpdateBuilder {
	return squirrel.Update(c.table)
}

func (c *customTkUserModel) DeleteBuilder() squirrel.DeleteBuilder {
	return squirrel.Delete(c.table)
}

func (c *customTkUserModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*TkUser, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp TkUser
	err = c.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

func (c *customTkUserModel) FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error) {

	query, values, err := sumBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64
	err = c.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (c *customTkUserModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

	query, values, err := countBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = c.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (c *customTkUserModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*TkUser, error) {

	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*TkUser
	err = c.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (c *customTkUserModel) FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*TkUser, error) {

	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := rowBuilder.Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*TkUser
	err = c.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (c *customTkUserModel) FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*TkUser, error) {

	if preMinId > 0 {
		rowBuilder = rowBuilder.Where(" id < ? ", preMinId)
	}

	query, values, err := rowBuilder.OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*TkUser
	err = c.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// 按照id升序分页查询数据，不支持排序
func (c *customTkUserModel) FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*TkUser, error) {

	if preMaxId > 0 {
		rowBuilder = rowBuilder.Where(" id > ? ", preMaxId)
	}

	query, values, err := rowBuilder.OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*TkUser
	err = c.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (c *customTkUserModel) InsertBatch(ctx context.Context, session sqlx.Session, insertBuilder squirrel.InsertBuilder) (sql.Result, error) {
	query, values, err := insertBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var ret sql.Result
	if session != nil {
		ret, err = session.ExecCtx(ctx, query, values...)
	} else {
		ret, err = c.conn.ExecCtx(ctx, query, values...)
	}

	return ret, err
}

func (c *customTkUserModel) UpdateBatch(ctx context.Context, session sqlx.Session, updateBuilder squirrel.UpdateBuilder) (sql.Result, error) {
	query, values, err := updateBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var ret sql.Result
	if session != nil {
		ret, err = session.ExecCtx(ctx, query, values...)
	} else {
		ret, err = c.conn.ExecCtx(ctx, query, values...)
	}

	return ret, err
}

func (c *customTkUserModel) DeleteBatch(ctx context.Context, session sqlx.Session, deleteBuilder squirrel.DeleteBuilder) (sql.Result, error) {
	query, values, err := deleteBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var ret sql.Result
	if session != nil {
		ret, err = session.ExecCtx(ctx, query, values...)
	} else {
		ret, err = c.conn.ExecCtx(ctx, query, values...)
	}

	return ret, err
}
