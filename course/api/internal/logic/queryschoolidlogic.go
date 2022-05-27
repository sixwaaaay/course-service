package logic

import (
	"context"
	"course/constants"
	"course/course/api/internal/errorx"
	"course/course/api/internal/svc"
	"course/course/api/internal/types"
	"course/school/rpc/types/school"
)

type QueryLogic func(req *types.QuerySchoolIdReq) (resp *types.BaseResponse, err error)

func NewQuerySchoolIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryLogic {
	return func(req *types.QuerySchoolIdReq) (resp *types.BaseResponse, err error) {
		rpcRes, err := svcCtx.SchoolRpc.QueryIdByName(ctx, &school.QueryIdByNameRequest{Name: req.Name})
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
}
