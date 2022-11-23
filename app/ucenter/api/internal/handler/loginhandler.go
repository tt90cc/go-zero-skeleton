package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"gomicro/app/ucenter/api/internal/logic"
	"gomicro/app/ucenter/api/internal/svc"
	"gomicro/app/ucenter/api/internal/types"
	"gomicro/common/response"
	"net/http"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.Response(w, resp, err)

	}
}
