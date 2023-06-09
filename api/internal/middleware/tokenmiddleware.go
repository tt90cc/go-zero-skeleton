package middleware

import (
	"context"
	"encoding/json"
	"gitee.com/qidianbox/common/ctxdata"
	"gitee.com/qidianbox/toker-engine-ucenter/rpc/ucenter"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"sync"
)

type TokenMiddleware struct {
	UcenterRpc ucenter.Ucenter
}

func NewTokenMiddleware(rpc ucenter.Ucenter) *TokenMiddleware {
	return &TokenMiddleware{
		UcenterRpc: rpc,
	}
}

func (m *TokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userInfo := ctxdata.GetUserInfoFromCtx(r.Context())

		reqBody := &ucenter.MenuPermissionReq{
			PlatformType: userInfo.PlatformType,
			UserId:       userInfo.Id,
			RoleId:       userInfo.Role,
			Path:         r.URL.Path,
		}

		var resp map[string]interface{}
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			if response, err := m.UcenterRpc.MenuPermission(r.Context(), reqBody); err != nil {
				logx.Infof("[DataPermission] get data permission failed. err:%v params:%+v", err, reqBody)
				resp = make(map[string]interface{})
				resp["data_permission"] = 3 // 个人
				resp["user_str"] = userInfo.Id
			} else {
				b, _ := json.Marshal(response)
				_ = json.Unmarshal(b, &resp)
			}

			wg.Done()
		}()

		// 验证用户状态
		var userinfoReply = new(ucenter.UserinfoReply)
		go func() {
			userinfoReply, _ = m.UcenterRpc.Userinfo(r.Context(), &ucenter.UserInfoReq{Id: userInfo.Id})
			wg.Done()
		}()

		wg.Wait()

		if userinfoReply.Status != 1 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ctxdata.CtxKeyPermission, resp)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
