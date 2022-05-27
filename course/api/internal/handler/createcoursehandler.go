package handler

import (
	"net/http"

	"course/course/api/internal/logic"
	"course/course/api/internal/svc"
	"course/course/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func createCourseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req types.CreateCourseReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		solver := logic.NewCreateLogic(r.Context(), svcCtx)
		resp, err := solver(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
