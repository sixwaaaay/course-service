package logic

import (
	"context"
	"course/constants"
	"course/course/api/internal/errorx"
	"course/course/api/internal/svc"
	"course/course/api/internal/types"
	"course/course/rpc/types/course"
)

type CheckLogic func(req *types.NameCheckReq) (resp *types.BaseResponse, err error)

func NewNameCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) CheckLogic {
	return func(req *types.NameCheckReq) (resp *types.BaseResponse, err error) {
		checkCourse, err := svcCtx.CourseRpc.CheckCourse(ctx, &course.CheckCourseRequest{
			Name: req.Name,
		})
		if err != nil {
			return nil, errorx.NewCodeError(constants.RpcErrCode, "failed to check course name")
		}
		if checkCourse.BaseResp.StatusCode != 0 {
			return nil, errorx.NewCodeError(checkCourse.BaseResp.StatusCode, checkCourse.BaseResp.StatusMessage)
		}
		return &types.BaseResponse{
			Status:  0,
			Message: "can use",
			Data:    nil,
		}, nil
	}
}
