package logic

import (
	"context"

	"gomicro/app/ucenter/rpc/internal/svc"
	"gomicro/app/ucenter/rpc/types/ucenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *ucenter.UserInfoReq) (*ucenter.UserInfoReply, error) {
	return &ucenter.UserInfoReply{Account: "allen"}, nil
}
