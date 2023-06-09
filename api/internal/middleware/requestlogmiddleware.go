package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/token"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type RequestLogMiddleware struct {
	AccessSecret string
}

func NewRequestLogMiddleware(accessSecret string) *RequestLogMiddleware {
	return &RequestLogMiddleware{
		AccessSecret: accessSecret,
	}
}

func (m *RequestLogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logx.WithContext(r.Context())

		bodyByte, _ := ioutil.ReadAll(r.Body)
		body := string(bodyByte)
		body = strings.ReplaceAll(body, " ", "")
		body = strings.ReplaceAll(body, "\n", "")

		var claimsByte []byte
		parser := token.NewTokenParser()
		if tok, err := parser.ParseToken(r, m.AccessSecret, ""); err == nil && tok.Valid {
			if claims, ok := tok.Claims.(jwt.MapClaims); ok {
				claimsByte, _ = json.Marshal(claims)
			}
		}

		requestPath := r.URL.Path
		if r.URL.RawQuery != "" {
			rawQuery, _ := url.PathUnescape(r.URL.RawQuery)
			requestPath = requestPath + "?" + rawQuery
		}

		logger.Infof("【请求参数】url:%s token:%s tokenClaims:%s body:%s", requestPath, r.Header.Get("Authorization"), string(claimsByte), body)
		r.Body = ioutil.NopCloser(bytes.NewReader(bodyByte))

		next(w, r)
	}
}
