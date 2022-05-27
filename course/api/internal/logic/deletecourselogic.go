package logic

import (
	"context"
	"course/constants"
	"course/course/api/internal/errorx"
	"course/course/rpc/types/course"

	"course/course/api/internal/svc"
	"course/course/api/internal/types"
)

type DeleteLogicFunc func(req *types.DeleteCourseReq) (resp *types.BaseResponse, err error)

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteLogicFunc {
	return func(req *types.DeleteCourseReq) (resp *types.BaseResponse, err error) {
		delResp, err := svcCtx.CourseRpc.DeleteCourse(ctx, &course.DeleteCourseRequest{
			Id: req.Id,
		})
		if err != nil {
			return nil, errorx.NewCodeError(constants.RpcErrCode, "fail")
		}
		if delResp.BaseResp.StatusCode != 0 {
			return nil, errorx.NewCodeError(delResp.BaseResp.StatusCode, delResp.BaseResp.StatusMessage)
		}
		return &types.BaseResponse{
			Status:  constants.SuccessCode,
			Message: "success",
			Data:    nil,
		}, nil
	}
}
