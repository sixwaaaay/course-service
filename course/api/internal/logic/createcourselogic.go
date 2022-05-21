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
		return nil, errorx.NewCodeError(constants.RpcErrCode, "fail")
	}
	if c.BaseResp.StatusCode != constants.SuccessCode {
		l.Logger.Errorf("create course failed, code: %d, msg: %s", c.BaseResp.StatusCode, c.BaseResp.StatusMessage)
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
