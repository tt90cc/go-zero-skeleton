package response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"gomicro/common/errorx"
	"net/http"
)

type Body struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// 统一封装成功响应值
func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		switch e := err.(type) {
		case *errorx.CodeError: // 业务输出错误
			body.Code = e.Code
			body.Message = e.Message
			body.Data = e.Data
		default: // 系统未知错误
			body.Code = errorx.ErrDEFAULT
			body.Message = errorx.ErrMSG[body.Code]
		}
	} else {
		body.Code = errorx.ErrOK
		body.Message = errorx.ErrMSG[body.Code]
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
