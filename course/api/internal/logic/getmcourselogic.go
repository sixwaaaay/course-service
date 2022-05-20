package logic

import (
	"context"
	"course/constants"
	"course/course/api/internal/pack"
	"course/course/api/internal/svc"
	"course/course/api/internal/types"
	"course/course/rpc/types/course"
	"course/school/rpc/types/school"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMCourseLogic {
	return &GetMCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMCourseLogic) GetMCourse(req *types.MCourseReq) (resp *types.BaseResponse, err error) {
	// covert pageSize and pageNo to offset and limit
	queryBySidResp, err := l.svcCtx.CourseRpc.QueryCourseBySid(l.ctx, &course.QueryCourseBySidRequest{
		Sid:    &req.Sid,
		Limit:  req.Size,
		Offset: (req.Page - 1) * req.Size,
	})
	if err != nil {
		return pack.BuildResp(constants.RPCInternalError, "fail to solve"), err
	}
	if queryBySidResp.BaseResp.StatusCode != constants.SuccessCode {
		return pack.BuildResp(queryBySidResp.BaseResp.StatusCode, queryBySidResp.BaseResp.StatusMessage), nil
	}
	schoolIds := pack.SchoolIds(queryBySidResp.Courses)
	schools, err := l.svcCtx.SchoolRpc.MGetSchool(l.ctx, &school.MGetSchoolRequest{SchoolIds: schoolIds})
	list := pack.ToCourseList(queryBySidResp.Courses)

	if err != nil {
		return pack.BuildResp(constants.RPCInternalError, "fail to solve"), err
	}
	if schools.BaseResp.StatusCode != 0 {
		l.Errorf("get school error: %s", schools.BaseResp.StatusMessage)
		return pack.BuildResp(schools.BaseResp.StatusCode, schools.BaseResp.StatusMessage), nil
	}
	schoolMap := pack.SchoolMap(schools.Schools)
	for i := 0; i < len(list); i++ {
		if s, ok := schoolMap[list[i].Sid]; ok {
			list[i].SName = s.Name
		}
	}
	resp = &types.BaseResponse{
		Status:  0,
		Message: "success",
		Data: &types.MCourseRes{
			Courses: list,
			Total:   queryBySidResp.Total,
		},
	}
	return
}
