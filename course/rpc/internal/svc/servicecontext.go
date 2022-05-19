package svc

import (
	"course/course/rpc/internal/config"
	"course/course/rpc/internal/db"
)

type ServiceContext struct {
	Config      config.Config
	CourseModel *db.CourseModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	model := db.NewCourseModel(c.DB.DataSource)

	return &ServiceContext{
		Config:      c,
		CourseModel: model,
	}
}
