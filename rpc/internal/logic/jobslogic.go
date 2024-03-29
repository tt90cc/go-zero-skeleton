package logic

import (
	"context"
	"fmt"
	"github.com/tt90cc/utils/globalkey"
	"github.com/zeromicro/go-zero/core/logx"
	"tt90.cc/ucenter/internal/svc"
)

type JobsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJobsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JobsLogic {
	return &JobsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JobsLogic) HelloWorld() {
	tryLock := l.svcCtx.TryLock(fmt.Sprintf(globalkey.RedisLock, "demo"), 10)
	if !tryLock {
		l.Logger.Info("get lock failed.")
		return
	}
	l.Logger.Info("Every second todo")
}
