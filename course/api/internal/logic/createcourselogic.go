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

type CreateCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCourseLogic {
	return &CreateCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCourseLogic) CreateCourse(req *types.CreateCourseReq) (resp *types.BaseResponse, err error) {
	c, err := l.svcCtx.CourseRpc.CreateCourse(l.ctx, &course.CreateCourseRequest{
		Name:   req.Name,
		Avatar: req.Avatar,
		Hours:  req.Hours,
		Sid:    req.Sid,
	})
	if err != nil {
		return pack.BuildResp(constants.RPCInternalError, "fail"), nil
	}
	if c.BaseResp.StatusCode != constants.SuccessCode {
		return &types.BaseResponse{
			Status:  c.BaseResp.StatusCode,
			Message: c.BaseResp.StatusMessage,
			Data:    nil,
		}, nil
	}
	return &types.BaseResponse{
		Status:  0,
		Message: "success",
		Data:    nil,
	}, nil
}
