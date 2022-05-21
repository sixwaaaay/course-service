package logic

import (
	"context"
	"course/constants"
	"course/course/rpc/internal/pack"
	"course/course/rpc/internal/svc"
	"course/course/rpc/types/course"

	"github.com/zeromicro/go-zero/core/logx"
)

type MGetCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMGetCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MGetCourseLogic {
	return &MGetCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MGetCourseLogic) MGetCourse(in *course.MGetCourseRequest) (*course.MGetCourseResponse, error) {
	mCourses, err := l.svcCtx.CourseModel.MGetCourses(l.ctx, in.CourseIds)
	if err != nil {
		return &course.MGetCourseResponse{
			Courses:  nil,
			BaseResp: pack.BuildResp(constants.DbErrCode, err.Error()),
		}, err
	}
	courses := make([]*course.Course, len(mCourses))
	for i, mCourse := range mCourses {
		courses[i] = &course.Course{
			Id:     mCourse.ID,
			Name:   mCourse.Name,
			Avatar: mCourse.Avatar,
			Hours:  mCourse.Hours,
			Sid:    mCourse.Sid,
		}
	}
	return &course.MGetCourseResponse{
		Courses:  courses,
		BaseResp: pack.BuildResp(constants.SuccessCode, "success"),
	}, nil
}
