// Code generated by goctl. DO NOT EDIT!
// Source: school.proto

package schoolservice

import (
	"context"

	"course/school/rpc/types/school"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BaseResp              = school.BaseResp
	MGetSchoolRequest     = school.MGetSchoolRequest
	MGetSchoolResponse    = school.MGetSchoolResponse
	QueryIdByNameRequest  = school.QueryIdByNameRequest
	QueryIdByNameResponse = school.QueryIdByNameResponse
	School                = school.School

	SchoolService interface {
		MGetSchool(ctx context.Context, in *MGetSchoolRequest, opts ...grpc.CallOption) (*MGetSchoolResponse, error)
		QueryIdByName(ctx context.Context, in *QueryIdByNameRequest, opts ...grpc.CallOption) (*QueryIdByNameResponse, error)
	}

	defaultSchoolService struct {
		cli zrpc.Client
	}
)

func NewSchoolService(cli zrpc.Client) SchoolService {
	return &defaultSchoolService{
		cli: cli,
	}
}

func (m *defaultSchoolService) MGetSchool(ctx context.Context, in *MGetSchoolRequest, opts ...grpc.CallOption) (*MGetSchoolResponse, error) {
	client := school.NewSchoolServiceClient(m.cli.Conn())
	return client.MGetSchool(ctx, in, opts...)
}

func (m *defaultSchoolService) QueryIdByName(ctx context.Context, in *QueryIdByNameRequest, opts ...grpc.CallOption) (*QueryIdByNameResponse, error) {
	client := school.NewSchoolServiceClient(m.cli.Conn())
	return client.QueryIdByName(ctx, in, opts...)
}