package logic

import (
	"context"
	"gomicro/app/ucenter/rpc/types/ucenter"
	"gomicro/common/errorx"

	"gomicro/app/ucenter/api/internal/svc"
	"gomicro/app/ucenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserinfoLogic {
	return &UserinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserinfoLogic) Userinfo() (resp *types.UserinfoReply, err error) {
	userInfo, err := l.svcCtx.UcenterRpc.UserInfo(l.ctx, &ucenter.UserInfoReq{Id: 1})
	if err != nil {
		return nil, errorx.NewCodeError(errorx.ErrDEFAULT)
	}
	return &types.UserinfoReply{
		Id:      userInfo.Id,
		Account: userInfo.Account,
	}, nil
}
