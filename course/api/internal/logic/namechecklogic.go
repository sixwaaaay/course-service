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

type NameCheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNameCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NameCheckLogic {
	return &NameCheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NameCheckLogic) NameCheck(req *types.NameCheckReq) (resp *types.BaseResponse, err error) {
	checkCourse, err := l.svcCtx.CourseRpc.CheckCourse(l.ctx, &course.CheckCourseRequest{
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
