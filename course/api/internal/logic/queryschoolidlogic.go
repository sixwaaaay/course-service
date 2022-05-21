package logic

import (
	"context"
	"course/constants"
	"course/course/api/internal/errorx"
	"course/school/rpc/types/school"

	"course/course/api/internal/svc"
	"course/course/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySchoolIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySchoolIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySchoolIdLogic {
	return &QuerySchoolIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuerySchoolIdLogic) QuerySchoolId(req *types.QuerySchoolIdReq) (resp *types.BaseResponse, err error) {
	rpcRes, err := l.svcCtx.SchoolRpc.QueryIdByName(l.ctx, &school.QueryIdByNameRequest{Name: req.Name})
	resp = &types.BaseResponse{}
	if err != nil {
		return nil, errorx.NewCodeError(constants.RpcErrCode, "fail")
	}
	if rpcRes.BaseResp.StatusCode != 0 {
		return nil, errorx.NewCodeError(rpcRes.BaseResp.StatusCode, rpcRes.BaseResp.StatusMessage)
	}
	resp.Status = constants.SuccessCode
	resp.Message = "success"
	resp.Data = rpcRes.SchoolId
	return
}
