package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/logx"

	"gateway/internal/config"
	"gateway/internal/handler"
	"gateway/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/gateway.yaml", "the config file")

func main() {
	flag.Parse()

	logx.MustSetup(logx.LogConf{
		Mode: "file",
		Path: "logs",
	})

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	logx.Infof("%+v", c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	logx.DisableStat()
	logx.Infof("Starting server at %s:%d...", c.Host, c.Port)
	server.Start()
}
