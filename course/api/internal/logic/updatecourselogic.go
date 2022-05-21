package logic

import (
	"context"
	"course/constants"
	"course/course/api/internal/errorx"
	"course/course/api/internal/svc"
	"course/course/api/internal/types"
	"course/course/rpc/types/course"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCourseLogic {
	return &UpdateCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCourseLogic) UpdateCourse(req *types.UpdateCourseReq) (resp *types.BaseResponse, err error) {
	reqReq := &course.UpdateCourseRequest{
		Id:     req.Id,
		Hours:  req.Hours,
		Sid:    req.Sid,
		Name:   req.Name,
		Avatar: req.Avatar,
	}

	updateCourse, err := l.svcCtx.CourseRpc.UpdateCourse(l.ctx, reqReq)
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
