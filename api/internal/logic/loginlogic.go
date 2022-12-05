package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"time"
	"tt90.cc/ucenter/api/internal/svc"
	"tt90.cc/ucenter/api/internal/types"
	"tt90.cc/ucenter/common/ctxdata"

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
		resp.AccessExpire = now + l.svcCtx.Config.Auth.AccessExpire
		resp.RefreshAfter = now + l.svcCtx.Config.Auth.AccessExpire/2
	}
	return
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[ctxdata.CtxKeyJwtUserId] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
