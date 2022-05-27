package logic

import (
	"context"
	"course/constants"
	"course/course/api/internal/errorx"
	"course/course/rpc/types/course"

	"course/course/api/internal/svc"
	"course/course/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogicFunc func(req *types.CreateCourseReq) (*types.BaseResponse, error)

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateLogicFunc {
	return func(req *types.CreateCourseReq) (*types.BaseResponse, error) {
		logger := logx.WithContext(ctx)
		c, err := svcCtx.CourseRpc.CreateCourse(ctx, &course.CreateCourseRequest{
			Name:   req.Name,
			Avatar: req.Avatar,
			Hours:  req.Hours,
			Sid:    req.Sid,
		})
		if err != nil {
			return nil, errorx.NewCodeError(constants.RpcErrCode, "fail")
		}
		if c.BaseResp.StatusCode != constants.SuccessCode {
			logger.Errorf("create course failed, code: %d, msg: %s", c.BaseResp.StatusCode, c.BaseResp.StatusMessage)
			return &types.BaseResponse{
				Status:  c.BaseResp.StatusCode,
				Message: "fail", // avoid expose internal error message like table name or column name
				Data:    nil,
			}, nil
		}
		return &types.BaseResponse{
			Status:  0,
			Message: "success",
			Data:    nil,
		}, nil
	}
}
