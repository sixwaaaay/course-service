package handler

import (
	"net/http"

	"course/course/api/internal/logic"
	"course/course/api/internal/svc"
	"course/course/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func deleteCourseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteCourseReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		resp, err := logic.NewDeleteLogic(r.Context(), svcCtx)(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
