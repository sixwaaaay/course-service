// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	schoolFieldNames          = builder.RawFieldNames(&School{})
	schoolRows                = strings.Join(schoolFieldNames, ",")
	schoolRowsExpectAutoSet   = strings.Join(stringx.Remove(schoolFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	schoolRowsWithPlaceHolder = strings.Join(stringx.Remove(schoolFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheSchoolIdPrefix   = "cache:school:id:"
	cacheSchoolNamePrefix = "cache:school:name:"
)

type (
	schoolModel interface {
		Insert(ctx context.Context, data *School) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*School, error)
		FindOneByName(ctx context.Context, name string) (*School, error)
		Update(ctx context.Context, data *School) error
		Delete(ctx context.Context, id int64) error
		MGetSchool(ctx context.Context, ids []int32) ([]*School, error)
	}

	defaultSchoolModel struct {
		sqlc.CachedConn
		table string
	}

	School struct {
		Id   int64  `db:"id"`
		Name string `db:"name"`
	}
)

func newSchoolModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultSchoolModel {
	return &defaultSchoolModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`s_school`", // 表名
	}
}

func (m *defaultSchoolModel) Insert(ctx context.Context, data *School) (sql.Result, error) {
	schoolIdKey := fmt.Sprintf("%s%v", cacheSchoolIdPrefix, data.Id)
	schoolNameKey := fmt.Sprintf("%s%v", cacheSchoolNamePrefix, data.Name)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?)", m.table, schoolRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name)
	}, schoolIdKey, schoolNameKey)
	return ret, err
}

func (m *defaultSchoolModel) FindOne(ctx context.Context, id int64) (*School, error) {
	schoolIdKey := fmt.Sprintf("%s%v", cacheSchoolIdPrefix, id)
	var resp School
	err := m.QueryRowCtx(ctx, &resp, schoolIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", schoolRows, m.table)
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

func (m *defaultSchoolModel) FindOneByName(ctx context.Context, name string) (*School, error) {
	schoolNameKey := fmt.Sprintf("%s%v", cacheSchoolNamePrefix, name)
	var resp School
	err := m.QueryRowIndexCtx(ctx, &resp, schoolNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", schoolRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, name); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSchoolModel) Update(ctx context.Context, data *School) error {
	schoolIdKey := fmt.Sprintf("%s%v", cacheSchoolIdPrefix, data.Id)
	schoolNameKey := fmt.Sprintf("%s%v", cacheSchoolNamePrefix, data.Name)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, schoolRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Name, data.Id)
	}, schoolIdKey, schoolNameKey)
	return err
}

func (m *defaultSchoolModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	schoolIdKey := fmt.Sprintf("%s%v", cacheSchoolIdPrefix, id)
	schoolNameKey := fmt.Sprintf("%s%v", cacheSchoolNamePrefix, data.Name)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, schoolIdKey, schoolNameKey)
	return err
}

func (m *defaultSchoolModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheSchoolIdPrefix, primary)
}

func (m *defaultSchoolModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", schoolRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSchoolModel) tableName() string {
	return m.table
}

func (m *defaultSchoolModel) MGetSchool(ctx context.Context, ids []int32) ([]*School, error) {
	// make a query string use ? as placeholder
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id CONCAT_WS(',', ?)", schoolRows, m.table)
	// prepare the query
	var resp []*School
	err := m.QueryRowsNoCacheCtx(ctx, resp, query, ids)
	return resp, err

	//return m.QueryRowCtx(ctx,)
}