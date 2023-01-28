package boot

import (
	"flag"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	g "tiktok/app/global"
	"tiktok/app/router"
	"time"
)

func ServerSetup() {
	config := g.Config.Server
	flag.Parse()
	readTimeout, err := time.ParseDuration(g.Config.Server.ReadTimeout)
	if err != nil {
		g.Logger.Errorf("parse duration err")
	}
	writeTimout, err := time.ParseDuration(g.Config.Server.ReadTimeout)
	if err != nil {
		g.Logger.Errorf("parse duration err")
	}
	h := server.Default(
		server.WithHostPorts(fmt.Sprintf(":%s", g.Config.Server.Port)),
		server.WithReadTimeout(readTimeout),
		server.WithWriteTimeout(writeTimout))
	router.InitRouter(h)
	h.Spin()
	g.Logger.Infof("server running on %s ...", config.Addr())
}
