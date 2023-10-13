package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"love_mall/user/internal/config"
	"love_mall/user/internal/handler"
	"love_mall/user/internal/model"
	"love_mall/user/internal/svc"
	"love_mall/utils"
)

var configFile = flag.String("f", "./etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	model.InitDb(c)

	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(utils.UnauthorizedCallback))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
