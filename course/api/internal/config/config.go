package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	CourseRpc zrpc.RpcClientConf
	SchoolRpc zrpc.RpcClientConf
	MinIO     struct {
		Endpoint  string
		AccessKey string
		SecretKey string
		Bucket    string
	}
}
