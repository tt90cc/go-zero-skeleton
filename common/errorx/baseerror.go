package errorx

import "github.com/spf13/cast"

type CodeError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (r *CodeError) Error() string {
	return r.Message
}

func NewCodeError(code int) error {
	return &CodeError{
		Code:    code,
		Message: MapErrMsg(code),
	}
}

func MapErrMsg(errcode int) string {
	if msg, ok := message[cast.ToUint32(errcode)]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

type CodeErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (r *CodeError) Info() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code:    r.Code,
		Message: r.Message,
		Data:    r.Data,
	}
}
