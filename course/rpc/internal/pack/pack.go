package pack

import (
	"course/course/rpc/types/course"
	"time"
)

func BuildResp(code int32, message string) *course.BaseResp {
	return &course.BaseResp{
		StatusCode:    code,
		StatusMessage: message,
		ServiceTime:   time.Now().Unix(),
	}
}
