package pack

import (
	"course/school/rpc/types/school"
	"time"
)

func BuildResp(code int32, message string) *school.BaseResp {
	return &school.BaseResp{
		StatusCode:    code,
		StatusMessage: message,
		ServiceTime:   time.Now().Unix(),
	}
}
