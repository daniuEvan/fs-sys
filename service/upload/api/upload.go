package main

import (
	"flag"
	"fmt"
	"fs-sys/common/authorization"
	"fs-sys/service/upload/api/internal/config"
	"fs-sys/service/upload/api/internal/handler"
	"fs-sys/service/upload/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/upload.yaml", "the config file")

func main() {
	flag.Parse()
	logx.DisableStat()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)

	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(authorization.UnauthorizedHandler()))
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
