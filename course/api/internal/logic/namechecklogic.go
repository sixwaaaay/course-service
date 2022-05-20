package logic

import (
	"context"
	"course/constants"
	"course/course/api/internal/pack"
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
		return pack.BuildResp(constants.RPCInternalError, "fail to solve"), nil
	}
	if checkCourse.BaseResp.StatusCode != 0 {
		return pack.BuildResp(checkCourse.BaseResp.StatusCode, checkCourse.BaseResp.StatusMessage), nil
	}
	return &types.BaseResponse{
		Status:  0,
		Message: "can use",
		Data:    nil,
	}, nil

}
