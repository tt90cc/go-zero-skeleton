package handler

import (
	"github.com/tt90cc/utils/response"
	"net/http"
	"tt90.cc/ucenter/internal/logic"
	"tt90.cc/ucenter/internal/svc"
)

func UserinfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewUserinfoLogic(r.Context(), svcCtx)
		resp, err := l.Userinfo()
		response.Response(w, resp, err)

	}
}
