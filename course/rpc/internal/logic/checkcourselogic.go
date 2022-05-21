package logic

import (
	"context"
	"course/constants"
	"course/course/rpc/internal/pack"
	"course/course/rpc/internal/svc"
	"course/course/rpc/types/course"
	"github.com/zeromicro/go-zero/core/logx"
)

type CheckCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckCourseLogic {
	return &CheckCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckCourseLogic) CheckCourse(in *course.CheckCourseRequest) (*course.CheckCourseResponse, error) {
	duplicated, err := l.svcCtx.CourseModel.QueryDuplicateName(l.ctx, in.Name)
	if err != nil {
		return &course.CheckCourseResponse{
			BaseResp: pack.BuildResp(constants.DbErrCode, err.Error()),
		}, nil
	}
	if duplicated {
		return &course.CheckCourseResponse{
			BaseResp: pack.BuildResp(constants.NameDuplicate, "duplicated name")}, nil
	} else {
		return &course.CheckCourseResponse{
			BaseResp: pack.BuildResp(constants.SuccessCode, "no duplicated name")}, nil
	}
}
