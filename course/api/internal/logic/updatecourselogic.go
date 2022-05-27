package logic

import (
	"context"
	"course/constants"
	"course/course/api/internal/errorx"
	"course/course/api/internal/svc"
	"course/course/api/internal/types"
	"course/course/rpc/types/course"
)

type UpdateLogic func(req *types.UpdateCourseReq) (resp *types.BaseResponse, err error)

func NewUpdateCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateLogic {
	return func(req *types.UpdateCourseReq) (resp *types.BaseResponse, err error) {
		reqReq := &course.UpdateCourseRequest{
			Id:     req.Id,
			Hours:  req.Hours,
			Sid:    req.Sid,
			Name:   req.Name,
			Avatar: req.Avatar,
		}

		updateCourse, err := svcCtx.CourseRpc.UpdateCourse(ctx, reqReq)
		if err != nil {
			return nil, errorx.NewCodeError(constants.RpcErrCode, "fail")
		}
		if updateCourse.BaseResp.StatusCode != 0 {
			return nil, errorx.NewCodeError(updateCourse.BaseResp.StatusCode, updateCourse.BaseResp.StatusMessage)
		}
		resp = &types.BaseResponse{
			Status:  updateCourse.BaseResp.StatusCode,
			Message: updateCourse.BaseResp.StatusMessage,
		}
		return

	}
}
