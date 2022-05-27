package logic

import (
	"context"
	"github.com/minio/minio-go/v7"
	"io"

	"course/course/api/internal/svc"
	"course/course/api/internal/types"
)

type DownloadLogic func(req *types.DownloadRequest) error

func NewDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext, writer io.Writer) DownloadLogic {
	return func(req *types.DownloadRequest) error {
		object, err := svcCtx.MinioCLi.GetObject(ctx, svcCtx.Bucket, req.File, minio.GetObjectOptions{})
		if err != nil {
			return err
		}
		defer object.Close()
		written, err := io.Copy(writer, object)
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
}
