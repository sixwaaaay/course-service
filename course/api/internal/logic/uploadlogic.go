package logic

import (
	"course/course/api/internal/svc"
	"course/course/api/internal/types"
	"github.com/minio/minio-go/v7"
	"net/http"
)

type UploadLogic func() (resp *types.BaseResponse, err error)

func NewUploadLogic(r *http.Request, svcCtx *svc.ServiceContext) UploadLogic {
	return func() (resp *types.BaseResponse, err error) {
		ctx := r.Context()
		file, handler, err := r.FormFile("myFile") // form name
		if err != nil {
			return nil, err
		}
		defer file.Close()
		info, err := svcCtx.MinioCLi.PutObject(ctx, svcCtx.Bucket, handler.Filename, file, handler.Size, minio.PutObjectOptions{})
		if err != nil {
			return nil, err
		}
		return &types.BaseResponse{
			Status:  0,
			Message: "upload success",
			Data:    info,
		}, nil

	}
}
