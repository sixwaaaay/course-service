package logic

import (
	"context"
	"course/constants"
	"course/school/rpc/internal/pack"
	"course/school/rpc/internal/svc"
	"course/school/rpc/types/school"
	"github.com/zeromicro/go-zero/core/logx"
)

type MGetSchoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMGetSchoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MGetSchoolLogic {
	return &MGetSchoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MGetSchoolLogic) MGetSchool(in *school.MGetSchoolRequest) (*school.MGetSchoolResponse, error) {

	getSchool, err := l.svcCtx.SchoolModel.MGetSchool(l.ctx, in.SchoolIds)
	if err != nil {
		return &school.MGetSchoolResponse{
			Schools:  nil,
			BaseResp: pack.BuildResp(constants.DbErrCode, err.Error()),
		}, err
	}

	// pack getSchool to MGetSchoolResponse
	schools := make([]*school.School, len(getSchool))
	for i, v := range getSchool {
		schools[i] = &school.School{
			Id:   v.ID,
			Name: v.Name,
		}
	}

	return &school.MGetSchoolResponse{
		Schools:  schools,
		BaseResp: pack.BuildResp(constants.SuccessCode, ""),
	}, nil
}
