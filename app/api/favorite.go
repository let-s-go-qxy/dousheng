package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jinzhu/copier"
	"strconv"
	"tiktok/app/global"
	"tiktok/app/internal/service"
)

// GetFavoriteList 获取Feed列表
func GetFavoriteList(c context.Context, ctx *app.RequestContext) {
	//1、token校验

	//2、请求处理
	userId := ctx.Query("user_id")
	uid, err := strconv.Atoi(userId)
	if err != nil {
		global.Logger.Error("用户ID错误")
	}
	videoList, videosFavoriteCount, videosAuthor := service.GetFavoriteList(uid)
	println(len(videoList))
	respVideoList := make([]Video, 0)
	for _, video := range videoList {
		respVideo := Video{}
		copier.Copy(&respVideo, &video)
		respVideo.FavoriteCount = videosFavoriteCount[int(video.Id)]

		user := User{}
		author := videosAuthor[int(video.Id)]
		copier.Copy(&user, &author)

		respVideo.Author = user
		respVideo.IsFavorite = true
		respVideoList = append(respVideoList, respVideo)
	}

	resp := FavoriteListResponse{StatusCode: consts.StatusOK, StatusMsg: "返回成功", VideoList: respVideoList}
	ctx.JSON(consts.StatusOK, resp)
}

// FavoriteAction 点赞和取消点赞操作
func FavoriteAction(c context.Context, ctx *app.RequestContext) {

}
