package logic

import (
	"context"
	"course/constants"
	"course/course/api/internal/pack"
	"course/course/api/internal/svc"
	"course/course/api/internal/types"
	"course/course/rpc/types/course"
	"course/school/rpc/types/school"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCourseLogic {
	return &GetCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCourseLogic) GetCourse(req *types.CourseReq) (resp *types.BaseResponse, err error) {
	mGetCourse, err := l.svcCtx.CourseRpc.MGetCourse(l.ctx, &course.MGetCourseRequest{
		CourseIds: []int32{req.Id},
	})
	if err != nil {
		return pack.BuildResp(constants.RPCInternalError, "fail to solve"), err
	}
	if mGetCourse.BaseResp.StatusCode != constants.SuccessCode {
		return pack.BuildResp(mGetCourse.BaseResp.StatusCode, mGetCourse.BaseResp.StatusMessage), nil
	}
	res := pack.ToCourse(mGetCourse.Courses[0])
	mGetSchool, err := l.svcCtx.SchoolRpc.MGetSchool(l.ctx, &school.MGetSchoolRequest{
		SchoolIds: []int32{res.Sid},
	})
	if err != nil {
		return pack.BuildResp(constants.RPCInternalError, "fail to solve"), err
	}
	if mGetSchool.BaseResp.StatusCode != constants.SuccessCode {
		return pack.BuildResp(mGetSchool.BaseResp.StatusCode, mGetSchool.BaseResp.StatusMessage), nil
	}
	res.SName = mGetSchool.Schools[0].Name

	resp = &types.BaseResponse{
		Status:  constants.SuccessCode,
		Message: "success",
		Data:    res,
	}
	return
}
