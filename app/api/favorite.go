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
	videoList, videoFavoriteCount := service.GetFavoriteList(uid)
	println(len(videoList))
	respVideoList := make([]Video, 0)
	for _, video := range videoList {
		respVideo := Video{}
		copier.Copy(&respVideo, &video)
		respVideo.FavoriteCount = videoFavoriteCount[int(video.Id)]
		respVideo.IsFavorite = true
		respVideoList = append(respVideoList, respVideo)
	}

	resp := FavoriteListResponse{StatusCode: consts.StatusOK, StatusMsg: "返回成功", VideoList: respVideoList}
	ctx.JSON(consts.StatusOK, resp)
}
