package handler

import (
	"net/http"

	"course/course/api/internal/logic"
	"course/course/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := logic.NewUploadLogic(r, svcCtx)()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
