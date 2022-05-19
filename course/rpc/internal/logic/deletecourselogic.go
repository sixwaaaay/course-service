package logic

import (
	"context"
	"course/constants"
	"course/course/rpc/internal/pack"
	"course/course/rpc/internal/svc"
	"course/course/rpc/types/course"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCourseLogic {
	return &DeleteCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCourseLogic) DeleteCourse(in *course.DeleteCourseRequest) (*course.DeleteCourseResponse, error) {

	if in.Id <= 0 {
		return &course.DeleteCourseResponse{
			BaseResp: pack.BuildResp(constants.InvalidParamsCode, "invalid params"),
		}, nil
	}

	err := l.svcCtx.CourseModel.DeleteCourse(l.ctx, in.Id)
	if err != nil {
		return &course.DeleteCourseResponse{
			BaseResp: pack.BuildResp(constants.DatabaseErrorCode, err.Error()),
		}, nil
	}

	return &course.DeleteCourseResponse{
		BaseResp: pack.BuildResp(constants.SuccessCode, "success"),
	}, nil
}
