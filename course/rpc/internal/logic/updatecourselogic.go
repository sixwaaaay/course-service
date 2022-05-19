package logic

import (
	"context"
	"course/constants"
	"course/course/rpc/internal/pack"
	"course/course/rpc/internal/svc"
	"course/course/rpc/types/course"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCourseLogic {
	return &UpdateCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCourseLogic) UpdateCourse(in *course.UpdateCourseRequest) (*course.UpdateCourseResponse, error) {
	// todo: 更新课程
	var (
		name   *string
		avatar *string
		hours  *int32
		sid    *int32
	)
	if in.Name != "" {
		name = &in.Name
	}
	if in.Avatar != "" {
		avatar = &in.Avatar
	}
	if in.Hours > 0 {
		hours = &in.Hours
	}
	if in.Sid > 0 {
		sid = &in.Sid
	}
	err := l.svcCtx.CourseModel.UpdateCourse(l.ctx, in.Id, name, avatar, sid, hours)
	if err != nil {
		return &course.UpdateCourseResponse{
			BaseResp: pack.BuildResp(constants.DatabaseErrorCode, err.Error()),
		}, err
	}
	return &course.UpdateCourseResponse{
		BaseResp: pack.BuildResp(constants.SuccessCode, "success"),
	}, nil
}
