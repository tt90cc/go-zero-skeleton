package jobs

import (
	"context"
	"github.com/robfig/cron/v3"
	"gomicro/app/ucenter/rpc/internal/logic"
	"gomicro/app/ucenter/rpc/internal/svc"
)

func RegisterJobs(serverCtx *svc.ServiceContext) {
	crontab(serverCtx)
	queue(serverCtx)
}

func crontab(serverCtx *svc.ServiceContext) {
	c := cron.New(cron.WithSeconds())

	c.AddFunc("* * * * * *", func() {
		logic.NewJobsLogic(context.TODO(), serverCtx).HelloWorld()
	})

	c.Start()
}

func queue(serverCtx *svc.ServiceContext) {
	// todo consumer
}
