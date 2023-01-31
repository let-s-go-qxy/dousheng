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
	publicGroup.POST("/user/register", api.UserRegister)
	publicGroup.POST("/user/login", api.UserLogin)
	loggedGroup.GET("/user", api.UserInfo)
	publicGroup.GET("/publish/action", api.GetFollowerList)
	loggedGroup.GET("/publish/list", api.PublishList)
	publicGroup.GET("/favorite/list", api.GetFavoriteList)
	publicGroup.GET("/comment/list", api.GetCommentList)       // 查看视频评论列表
	loggedGroup.POST("/comment/action", api.PostCommentAction) // 修改视频评论
	loggedGroup.POST("/relation/action", api.PublishVideo)
	loggedGroup.GET("/relation/follow/list", api.GetFollowerList)
	loggedGroup.GET("/relation/follower/list", api.GetFollowerList)
	loggedGroup.GET("/relation/friend/list", api.GetFollowerList)
	publicGroup.GET("/message/chat", api.GetFollowerList)
	publicGroup.GET("/message/action", api.GetFollowerList)

	// 路由注册成功log
	g.Logger.Infof("initialize routers successfully")
}

// InitGroup 生成路由组，path为公共前缀
func InitGroup(h *server.Hertz, path string, handlers ...app.HandlerFunc) *route.RouterGroup {
	return h.Group("/douyin"+path, handlers...)
}
