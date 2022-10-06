package main

import (
	"flag"
	dtmLogger "github.com/dtm-labs/logger"
	"github.com/zeromicro/go-zero/core/logx"
	"seckill/internal/logic/lua"
	"seckill/service/seckill"

	"seckill/internal/config"
	"seckill/internal/server"
	"seckill/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/dtm-labs/driver-gozero"
)

var configFile = flag.String("f", "etc/seckill.yaml", "the config file")

type nopWriter struct{}

func (*nopWriter) Debugf(format string, args ...interface{}) {

}
func (*nopWriter) Infof(format string, args ...interface{}) {

}
func (*nopWriter) Warnf(format string, args ...interface{}) {

}
func (*nopWriter) Errorf(format string, args ...interface{}) {

}

func main() {
	flag.Parse()

	logx.MustSetup(logx.LogConf{
		Mode: "file",
		Path: "logs",
	})

	dtmLogger.InitLog("error")
	nw := nopWriter{}
	dtmLogger.WithLogger(&nw)

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	logx.Infof("%+v", c)
	ctx := svc.NewServiceContext(c)

	lua.LoadLuaScript(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		seckill.RegisterSeckillServer(grpcServer, server.NewSeckillServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	logx.Infof("Starting rpc server at %s...", c.ListenOn)
	s.Start()
}
