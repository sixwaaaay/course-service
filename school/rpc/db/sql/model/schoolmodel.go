package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SchoolModel = (*customSchoolModel)(nil)

type (
	// SchoolModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSchoolModel.
	SchoolModel interface {
		schoolModel
		//MGetSchool(ctx context.Context, ids []int32) (interface{}, interface{})
	}

	customSchoolModel struct {
		*defaultSchoolModel
	}
)

// NewSchoolModel returns a model for the database table.
func NewSchoolModel(conn sqlx.SqlConn, c cache.CacheConf) SchoolModel {
	return &customSchoolModel{
		defaultSchoolModel: newSchoolModel(conn, c),
	}
}
