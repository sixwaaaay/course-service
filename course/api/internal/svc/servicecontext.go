package svc

import (
	"course/course/api/internal/config"
	"course/course/rpc/courseservice"
	"course/school/rpc/schoolservice"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	CourseRpc courseservice.CourseService
	SchoolRpc schoolservice.SchoolService
	MinioCLi  *minio.Client
	Bucket    string
}

func NewServiceContext(c config.Config) *ServiceContext {
	client, err := minio.New(c.MinIO.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(c.MinIO.AccessKey, c.MinIO.SecretKey, ""),
	})
	if err != nil {
		panic(err.Error())
	}
	return &ServiceContext{
		Config:    c,
		CourseRpc: courseservice.NewCourseService(zrpc.MustNewClient(c.CourseRpc)),
		SchoolRpc: schoolservice.NewSchoolService(zrpc.MustNewClient(c.SchoolRpc)),
		MinioCLi:  client,
		Bucket:    c.MinIO.Bucket,
	}
}
