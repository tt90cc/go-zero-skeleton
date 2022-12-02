package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/threading"
	"tt90.cc/ucenter/rpc/internal/config"
	"tt90.cc/ucenter/rpc/internal/jobs"
	"tt90.cc/ucenter/rpc/internal/server"
	"tt90.cc/ucenter/rpc/internal/svc"
	"tt90.cc/ucenter/rpc/types/ucenter"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/ucenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewUcenterServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		ucenter.RegisterUcenterServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	threading.GoSafe(func() {
		jobs.RegisterJobs(ctx)
	})

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
