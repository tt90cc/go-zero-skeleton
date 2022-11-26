package response

import (
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"tt90.cc/ucenter/common/errorx"
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
		err = errors.Cause(err)
		switch e := err.(type) {
		case *errorx.CodeError: // 业务输出错误
			body.Code = e.Code
			body.Message = e.Message
			body.Data = e.Data
		default: // 系统未知错误
			body.Code = errorx.ERR_DEFAULT
			body.Message = errorx.MapErrMsg(body.Code)
		}
	} else {
		body.Code = errorx.OK
		body.Message = errorx.MapErrMsg(body.Code)
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
