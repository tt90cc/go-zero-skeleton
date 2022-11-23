package handler

import (
	"gomicro/app/ucenter/api/internal/logic"
	"gomicro/app/ucenter/api/internal/svc"
	"gomicro/common/response"
	"net/http"
)

func UserinfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewUserinfoLogic(r.Context(), svcCtx)
		resp, err := l.Userinfo()
		response.Response(w, resp, err)

	}
}
