package logic

import (
	"context"
	"course/constants"
	"course/course/rpc/internal/db"
	"course/course/rpc/internal/pack"
	"course/course/rpc/internal/svc"
	"course/course/rpc/types/course"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCourseLogic {
	return &CreateCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCourseLogic) CreateCourse(in *course.CreateCourseRequest) (*course.CreateCourseResponse, error) {
	err := l.svcCtx.CourseModel.CreateCourse(l.ctx,
		&db.Course{
			ID:     0,
			Name:   in.Name,
			Avatar: in.Avatar,
			Hours:  in.Hours,
			Sid:    in.Sid,
		},
	)
	if err != nil {
		return &course.CreateCourseResponse{
			BaseResp: pack.BuildResp(constants.DatabaseErrorCode, err.Error()),
		}, nil
	}
	return &course.CreateCourseResponse{
		BaseResp: pack.BuildResp(constants.SuccessCode, "success"),
	}, nil
}
