package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		// 示例
		// DB:
		//  DataSource: root:PXDN93VRKUm8TeE7@tcp(mysql:3306)/looklook_order?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai
		DataSource string
	}
}
