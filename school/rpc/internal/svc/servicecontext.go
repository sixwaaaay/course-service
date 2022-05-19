package svc

import (
	"course/school/rpc/db"
	"course/school/rpc/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	SchoolModel *db.SchoolModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		SchoolModel: db.NewSchoolModel(c.DB.DataSource),
	}
}
