package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"order/internal/config"
	"order/internal/logic"
	"order/internal/server"
	"order/internal/svc"
	"order/service/order"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/dtm-labs/driver-gozero"
	_ "net/http/pprof"
)

var configFile = flag.String("f", "etc/order.yaml", "the config file")

func main() {
	go http.ListenAndServe(":18083", nil)

	flag.Parse()

	logx.MustSetup(logx.LogConf{
		Mode: "file",
		Path: "logs",
	})

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	logx.Infof("%+v", c)
	ctx := svc.NewServiceContext(c)

	go logic.Consume(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		order.RegisterOrderServer(grpcServer, server.NewOrderServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	logx.Infof("Starting rpc server at %s...", c.ListenOn)
	s.Start()
}
