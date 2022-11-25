package handler

import (
	"net/http"
	"tt90.cc/ucenter/api/internal/logic"
	"tt90.cc/ucenter/api/internal/svc"
	"tt90.cc/ucenter/common/response"
)

func UserinfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewUserinfoLogic(r.Context(), svcCtx)
		resp, err := l.Userinfo()
		response.Response(w, resp, err)

	}
}
