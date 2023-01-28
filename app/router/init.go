package router

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/route"
	"tiktok/app/api"
	g "tiktok/app/global"
	"tiktok/app/internal/middleware"
)

// InitRouter 初始化路由
func InitRouter(h *server.Hertz) {

	// 初始化全局中间件
	middleware.InitMiddleWareForDefault(h)

	// 该路由组无需使用token中间件鉴权
	publicGroup := InitGroup(h, "")

	//privateGroup := InitGroup(h, "", middleware.Jwt())

	// 路由配置，跟上单独中间件
	publicGroup.GET("/relation/follow/list", api.GetFollowerList)
	//publicGroup.GET("/xxx",  api.GetXxx)
	//publicGroup.GET("/xxx",  api.GetXxx)

	g.Logger.Infof("initialize routers successfully")
}

// InitGroup 生成路由组，path为公共前缀
func InitGroup(h *server.Hertz, path string, handlers ...app.HandlerFunc) *route.RouterGroup {
	return h.Group("/douyin"+path, handlers...)
}
