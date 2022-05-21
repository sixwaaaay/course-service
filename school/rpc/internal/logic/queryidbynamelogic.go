package logic

import (
	"context"
	"course/constants"
	"course/school/rpc/internal/pack"

	"course/school/rpc/internal/svc"
	"course/school/rpc/types/school"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryIdByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryIdByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryIdByNameLogic {
	return &QueryIdByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryIdByNameLogic) QueryIdByName(in *school.QueryIdByNameRequest) (*school.QueryIdByNameResponse, error) {
	id, err := l.svcCtx.SchoolModel.GetIdByName(l.ctx, in.Name)
	if err != nil {
		return &school.QueryIdByNameResponse{
			SchoolId: 0,
			BaseResp: pack.BuildResp(constants.DbErrCode, "database error"),
		}, nil
	}
	return &school.QueryIdByNameResponse{
		SchoolId: id,
		BaseResp: pack.BuildResp(constants.SuccessCode, "success"),
	}, nil
}
