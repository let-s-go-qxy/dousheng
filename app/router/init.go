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
	// 该路由组需要使用token鉴权中间件
	loggedGroup := InitGroup(h, "", middleware.Jwt())

	// 路由配置，跟上单独中间件 注意看好请求方法和是否需要登录
	publicGroup.GET("/feed", api.GetFeedList)
	loggedGroup.GET("/favorite/action", api.GetFollowerList)
	publicGroup.GET("/user/register", api.GetFollowerList)
	publicGroup.GET("/user/login", api.GetFollowerList)
	publicGroup.GET("/user", api.GetUserInfo)
	publicGroup.GET("/publish/action", api.GetFollowerList)
	publicGroup.GET("/publish/list", api.GetFollowerList)
	publicGroup.GET("/favorite/list", api.GetFollowerList)
	publicGroup.GET("/comment/list", api.GetFollowerList)
	publicGroup.GET("/comment/action", api.GetFollowerList)
	loggedGroup.POST("/relation/action", api.PublishVideo)
	publicGroup.GET("/relation/follow/list", api.GetFollowerList)
	publicGroup.GET("/relation/follower/list", api.GetFollowerList)
	publicGroup.GET("/relation/friend/list", api.GetFollowerList)
	publicGroup.GET("/message/chat", api.GetFollowerList)
	publicGroup.GET("/message/action", api.GetFollowerList)

	// 路由注册成功log
	g.Logger.Infof("initialize routers successfully")
}

// InitGroup 生成路由组，path为公共前缀
func InitGroup(h *server.Hertz, path string, handlers ...app.HandlerFunc) *route.RouterGroup {
	return h.Group("/douyin"+path, handlers...)
}
