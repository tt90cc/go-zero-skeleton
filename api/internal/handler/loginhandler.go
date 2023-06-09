package handler

import (
	"github.com/tt90cc/utils/errorx"
	"github.com/tt90cc/utils/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"tt90.cc/ucenter/internal/logic"
	"tt90.cc/ucenter/internal/svc"
	"tt90.cc/ucenter/internal/types"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, errorx.NewCodeError(errorx.ERR_PARAMS, err.Error()))
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.Response(w, resp, err)

	}
}
