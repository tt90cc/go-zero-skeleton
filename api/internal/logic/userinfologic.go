package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"tt90.cc/ucenter/api/internal/svc"
	"tt90.cc/ucenter/api/internal/types"
	"tt90.cc/ucenter/common/errorx"
	"tt90.cc/ucenter/rpc/types/ucenter"

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
		return nil, errors.Wrapf(errorx.NewCodeError(errorx.ERR_DEFAULT), "请求UcenterRpc失败. id:%d,err:%v", 1, err)
	}
	_ = copier.Copy(&resp, userInfo)
	return resp, nil
}
