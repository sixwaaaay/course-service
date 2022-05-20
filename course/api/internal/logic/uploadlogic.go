package logic

import (
	"context"
	"course/course/api/internal/svc"
	"course/course/api/internal/types"
	"github.com/minio/minio-go/v7"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadLogic(r *http.Request, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(r.Context()),
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload() (resp *types.BaseResponse, err error) {
	file, handler, err := l.r.FormFile("myFile") // form name
	if err != nil {
		return nil, err
	}
	defer file.Close()
	info, err := l.svcCtx.MinioCLi.PutObject(l.ctx, l.svcCtx.Bucket, handler.Filename, file, handler.Size, minio.PutObjectOptions{})
	if err != nil {
		return nil, err
	}
	return &types.BaseResponse{
		Status:  0,
		Message: "upload success",
		Data:    info,
	}, nil
}
