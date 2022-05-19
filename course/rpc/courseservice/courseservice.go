// Code generated by goctl. DO NOT EDIT!
// Source: course.proto

package courseservice

import (
	"context"

	"course/course/rpc/types/course"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BaseResp                 = course.BaseResp
	CheckCourseRequest       = course.CheckCourseRequest
	CheckCourseResponse      = course.CheckCourseResponse
	Course                   = course.Course
	CreateCourseRequest      = course.CreateCourseRequest
	CreateCourseResponse     = course.CreateCourseResponse
	DeleteCourseRequest      = course.DeleteCourseRequest
	DeleteCourseResponse     = course.DeleteCourseResponse
	MGetCourseRequest        = course.MGetCourseRequest
	MGetCourseResponse       = course.MGetCourseResponse
	QueryCourseBySidRequest  = course.QueryCourseBySidRequest
	QueryCourseBySidResponse = course.QueryCourseBySidResponse
	UpdateCourseRequest      = course.UpdateCourseRequest
	UpdateCourseResponse     = course.UpdateCourseResponse

	CourseService interface {
		CreateCourse(ctx context.Context, in *CreateCourseRequest, opts ...grpc.CallOption) (*CreateCourseResponse, error)
		DeleteCourse(ctx context.Context, in *DeleteCourseRequest, opts ...grpc.CallOption) (*DeleteCourseResponse, error)
		UpdateCourse(ctx context.Context, in *UpdateCourseRequest, opts ...grpc.CallOption) (*UpdateCourseResponse, error)
		MGetCourse(ctx context.Context, in *MGetCourseRequest, opts ...grpc.CallOption) (*MGetCourseResponse, error)
		CheckCourse(ctx context.Context, in *CheckCourseRequest, opts ...grpc.CallOption) (*CheckCourseResponse, error)
		QueryCourseBySid(ctx context.Context, in *QueryCourseBySidRequest, opts ...grpc.CallOption) (*QueryCourseBySidResponse, error)
	}

	defaultCourseService struct {
		cli zrpc.Client
	}
)

func NewCourseService(cli zrpc.Client) CourseService {
	return &defaultCourseService{
		cli: cli,
	}
}

func (m *defaultCourseService) CreateCourse(ctx context.Context, in *CreateCourseRequest, opts ...grpc.CallOption) (*CreateCourseResponse, error) {
	client := course.NewCourseServiceClient(m.cli.Conn())
	return client.CreateCourse(ctx, in, opts...)
}

func (m *defaultCourseService) DeleteCourse(ctx context.Context, in *DeleteCourseRequest, opts ...grpc.CallOption) (*DeleteCourseResponse, error) {
	client := course.NewCourseServiceClient(m.cli.Conn())
	return client.DeleteCourse(ctx, in, opts...)
}

func (m *defaultCourseService) UpdateCourse(ctx context.Context, in *UpdateCourseRequest, opts ...grpc.CallOption) (*UpdateCourseResponse, error) {
	client := course.NewCourseServiceClient(m.cli.Conn())
	return client.UpdateCourse(ctx, in, opts...)
}

func (m *defaultCourseService) MGetCourse(ctx context.Context, in *MGetCourseRequest, opts ...grpc.CallOption) (*MGetCourseResponse, error) {
	client := course.NewCourseServiceClient(m.cli.Conn())
	return client.MGetCourse(ctx, in, opts...)
}

func (m *defaultCourseService) CheckCourse(ctx context.Context, in *CheckCourseRequest, opts ...grpc.CallOption) (*CheckCourseResponse, error) {
	client := course.NewCourseServiceClient(m.cli.Conn())
	return client.CheckCourse(ctx, in, opts...)
}

func (m *defaultCourseService) QueryCourseBySid(ctx context.Context, in *QueryCourseBySidRequest, opts ...grpc.CallOption) (*QueryCourseBySidResponse, error) {
	client := course.NewCourseServiceClient(m.cli.Conn())
	return client.QueryCourseBySid(ctx, in, opts...)
}