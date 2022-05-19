package logic

import (
	"context"
	"course/constants"
	"course/course/rpc/internal/pack"
	"time"

	"course/course/rpc/internal/svc"
	"course/course/rpc/types/course"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryCourseBySidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCourseBySidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCourseBySidLogic {
	return &QueryCourseBySidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryCourseBySidLogic) QueryCourseBySid(in *course.QueryCourseBySidRequest) (*course.QueryCourseBySidResponse, error) {

	courses, total, err := l.svcCtx.CourseModel.QueryCourseBySid(l.ctx, in.Sid, int(in.Limit), int(in.Offset))
	if err != nil {
		return &course.QueryCourseBySidResponse{
			Courses:  nil,
			Total:    0,
			BaseResp: pack.BuildResp(constants.DatabaseErrorCode, err.Error()),
		}, err
	}
	// covert to response type
	var coursesResp []*course.Course
	for _, v := range courses {
		coursesResp = append(coursesResp, &course.Course{
			Id:     v.ID,
			Name:   v.Name,
			Sid:    v.Sid,
			Hours:  v.Hours,
			Avatar: v.Avatar,
		})
	}

	return &course.QueryCourseBySidResponse{
		Courses: coursesResp,
		Total:   int32(total),
		BaseResp: &course.BaseResp{
			StatusCode:    0,
			StatusMessage: "success",
			ServiceTime:   time.Now().Unix(),
		},
	}, nil
}
