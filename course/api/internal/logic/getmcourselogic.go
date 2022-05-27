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

type MultiGetLogic func(req *types.MCourseReq) (resp *types.BaseResponse, err error)

func NewMultiGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) MultiGetLogic {
	return func(req *types.MCourseReq) (resp *types.BaseResponse, err error) {
		queryBySidResp, err := svcCtx.CourseRpc.QueryCourseBySid(ctx, &course.QueryCourseBySidRequest{
			Sid:    &req.Sid,
			Limit:  req.Size,
			Offset: (req.Page - 1) * req.Size,
		})
		if err != nil {
			return nil, errorx.NewCodeError(constants.RpcErrCode, "fail")
		}
		if queryBySidResp.BaseResp.StatusCode != constants.SuccessCode {
			return nil, errorx.NewCodeError(queryBySidResp.BaseResp.StatusCode, queryBySidResp.BaseResp.StatusMessage)
		}
		schoolIds := pack.SchoolIds(queryBySidResp.Courses)
		schools, err := svcCtx.SchoolRpc.MGetSchool(ctx, &school.MGetSchoolRequest{SchoolIds: schoolIds})
		list := pack.ToCourseList(queryBySidResp.Courses)

		if err != nil {
			return nil, errorx.NewCodeError(constants.RpcErrCode, "fail")
		}
		if schools.BaseResp.StatusCode != 0 {
			return nil, errorx.NewCodeError(schools.BaseResp.StatusCode, schools.BaseResp.StatusMessage)
		}
		schoolMap := pack.SchoolMap(schools.Schools)
		for i := 0; i < len(list); i++ {
			if s, ok := schoolMap[list[i].Sid]; ok {
				list[i].SName = s.Name
			}
		}
		resp = &types.BaseResponse{
			Status:  0,
			Message: "success",
			Data: &types.MCourseRes{
				Courses: list,
				Total:   queryBySidResp.Total,
			},
		}
		return
	}
}
