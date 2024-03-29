// Code generated by goctl. DO NOT EDIT!
// Source: main.proto

package server

import (
	"context"

	"tt90.cc/ucenter/internal/logic"
	"tt90.cc/ucenter/internal/svc"
	"tt90.cc/ucenter/types/ucenter"
)

type UcenterServer struct {
	svcCtx *svc.ServiceContext
	ucenter.UnimplementedUcenterServer
}

func NewUcenterServer(svcCtx *svc.ServiceContext) *UcenterServer {
	return &UcenterServer{
		svcCtx: svcCtx,
	}
}

func (s *UcenterServer) UserInfo(ctx context.Context, in *ucenter.UserInfoReq) (*ucenter.UserInfoReply, error) {
	l := logic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}
