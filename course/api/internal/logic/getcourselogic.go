package logic

import (
	"context"
	"course/constants"
	"course/course/api/internal/errorx"
	"course/course/api/internal/pack"
	"course/course/api/internal/svc"
	"course/course/api/internal/types"
	"course/course/rpc/types/course"
	"course/school/rpc/types/school"
)

type GetCourseLogicFunc func(req *types.CourseReq) (resp *types.BaseResponse, err error)

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetCourseLogicFunc {
	return func(req *types.CourseReq) (resp *types.BaseResponse, err error) {
		//logger := logx.WithContext(ctx)
		mGetCourse, err := svcCtx.CourseRpc.MGetCourse(ctx, &course.MGetCourseRequest{
			CourseIds: []int32{req.Id},
		})
		if err != nil {
			return nil, errorx.NewCodeError(constants.RpcErrCode, "fail")
		}
		if mGetCourse.BaseResp.StatusCode != constants.SuccessCode {
			return pack.BuildResp(mGetCourse.BaseResp.StatusCode, mGetCourse.BaseResp.StatusMessage), nil
		}
		res := pack.ToCourse(mGetCourse.Courses[0])
		mGetSchool, err := svcCtx.SchoolRpc.MGetSchool(ctx, &school.MGetSchoolRequest{
			SchoolIds: []int32{res.Sid},
		})
		if err != nil {
			return nil, errorx.NewCodeError(constants.RpcErrCode, "fail")
		}
		if mGetSchool.BaseResp.StatusCode != constants.SuccessCode {
			return nil, errorx.NewCodeError(mGetSchool.BaseResp.StatusCode, mGetSchool.BaseResp.StatusMessage)
		}
		res.SName = mGetSchool.Schools[0].Name

		resp = &types.BaseResponse{
			Status:  constants.SuccessCode,
			Message: "success",
			Data:    res,
		}
		return
	}
}
