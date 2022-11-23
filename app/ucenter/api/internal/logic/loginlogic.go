package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"gomicro/app/ucenter/api/internal/svc"
	"gomicro/app/ucenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginReply, err error) {
	resp = &types.LoginReply{
		Id:      1,
		Account: "Allen",
	}
	now := time.Now().Unix()
	if token, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, 1); err != nil {
		l.Logger.Errorf("[Login] getJwtToken failed. userid:%d, err:%+v", 1, err)
	} else {
		resp.AccessToken = token
	}
	return
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
