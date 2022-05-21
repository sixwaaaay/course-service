package errorx

import "course/course/api/internal/types"

const defaultCode = 1001

type CodeError struct {
	Code    int32  `json:"code"`
	Message string `json:"msg"`
}

func NewCodeError(code int32, msg string) error {
	return &CodeError{Code: code, Message: msg}
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}

func (e *CodeError) Error() string {
	return e.Message
}

func (e *CodeError) Data() *types.BaseResponse {
	return &types.BaseResponse{
		Status:  e.Code,
		Message: e.Message,
	}
}
