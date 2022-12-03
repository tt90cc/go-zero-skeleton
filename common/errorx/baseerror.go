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

func NewCodeError(code int, message ...string) error {
	msg := MapErrMsg(code)
	if len(message) > 0 {
		msg = message[0]
	}

	return &CodeError{
		Code:    code,
		Message: msg,
	}
}

func MapErrMsg(errcode int) string {
	if msg, ok := message[cast.ToUint32(errcode)]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

type codeErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (r *CodeError) Info() *codeErrorResponse {
	return &codeErrorResponse{
		Code:    r.Code,
		Message: r.Message,
		Data:    r.Data,
	}
}
