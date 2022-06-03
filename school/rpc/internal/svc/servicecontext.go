package svc

import (
	"course/school/rpc/db/sql/model"
	"course/school/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	SchoolModel model.SchoolModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:      c,
		SchoolModel: model.NewSchoolModel(sqlConn, c.Cache),
	}
}
