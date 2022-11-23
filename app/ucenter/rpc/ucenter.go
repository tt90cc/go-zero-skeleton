package main

import (
	"flag"
	"fmt"
	"gomicro/app/ucenter/rpc/internal/jobs"

	"gomicro/app/ucenter/rpc/internal/config"
	"gomicro/app/ucenter/rpc/internal/server"
	"gomicro/app/ucenter/rpc/internal/svc"
	"gomicro/app/ucenter/rpc/types/ucenter"

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

	jobs.RegisterJobs(ctx)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
