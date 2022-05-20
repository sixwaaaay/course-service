package logic

import (
	"context"
	"github.com/minio/minio-go/v7"
	"io"

	"course/course/api/internal/svc"
	"course/course/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	writer io.Writer
}

func NewDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext, writer io.Writer) *DownloadLogic {
	return &DownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		writer: writer,
	}
}

func (l *DownloadLogic) Download(req *types.DownloadRequest) error {
	object, err := l.svcCtx.MinioCLi.GetObject(l.ctx, l.svcCtx.Bucket, req.File, minio.GetObjectOptions{})
	if err != nil {
		return err
	}
	defer object.Close()
	written, err := io.Copy(l.writer, object)
	if err != nil {
		return err
	}
	stat, err := object.Stat()
	if err != nil {
		return err
	}
	if written != stat.Size {
		return io.ErrClosedPipe
	}
	return nil
}
