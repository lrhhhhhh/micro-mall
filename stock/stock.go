package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/logx"
	"stock/internal/config"
	"stock/internal/logic"
	"stock/internal/server"
	"stock/internal/svc"
	"stock/service/stock"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/dtm-labs/driver-gozero"
)

var configFile = flag.String("f", "etc/stock.yaml", "the config file")

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

	go logic.ConsumeDeductMessage(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		stock.RegisterStockServer(grpcServer, server.NewStockServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	logx.Infof("Starting rpc server at %s...", c.ListenOn)
	s.Start()
}
