package logic

import (
	"context"
	"course/constants"
	"course/course/api/internal/pack"
	"course/course/rpc/types/course"

	"course/course/api/internal/svc"
	"course/course/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCourseLogic {
	return &DeleteCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCourseLogic) DeleteCourse(req *types.DeleteCourseReq) (resp *types.BaseResponse, err error) {
	delResp, err := l.svcCtx.CourseRpc.DeleteCourse(l.ctx, &course.DeleteCourseRequest{
		Id: req.Id,
	})
	if err != nil {
		return pack.BuildResp(constants.RPCInternalError, "fail to solve"), err
	}
	if delResp.BaseResp.StatusCode != 0 {
		return pack.BuildResp(delResp.BaseResp.StatusCode, delResp.BaseResp.StatusMessage), nil
	}
	return &types.BaseResponse{
		Status:  constants.SuccessCode,
		Message: "success",
		Data:    nil,
	}, nil
}
