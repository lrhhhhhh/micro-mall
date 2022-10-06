package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/logx"
	"user/internal/config"
	"user/internal/server"
	"user/internal/svc"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	logx.MustSetup(logx.LogConf{
		Mode: "file",
		Path: "logs",
	})

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	logx.Infof("%+v", c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	logx.DisableStat()
	logx.Infof("Starting rpc server at %s...", c.ListenOn)
	s.Start()
}
