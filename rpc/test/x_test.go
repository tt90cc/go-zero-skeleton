package test

import (
	"context"
	"github.com/zeromicro/go-zero/core/conf"
	"testing"
	"tt90.cc/ucenter/internal/config"
	"tt90.cc/ucenter/internal/svc"
)

var ctx context.Context
var svcCtx *svc.ServiceContext

func init() {
	ctx = context.TODO()
	var c config.Config
	conf.MustLoad("etc/ucenter.yaml", &c)
	svcCtx = svc.NewServiceContext(c)
}

func TestHello(t *testing.T) {
	t.Log("Hello world.")
}
