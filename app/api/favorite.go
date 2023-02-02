package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jinzhu/copier"
	"strconv"
	"tiktok/app/global"
	"tiktok/app/internal/service"
	"tiktok/utils/msg"
)

// GetFavoriteList 获取喜爱视频列表
func GetFavoriteList(c context.Context, ctx *app.RequestContext) {
	userId := ctx.Query("user_id")
	uid, err := strconv.Atoi(userId)
	if err != nil {
		global.Logger.Error("用户ID错误")
	}
	videoList := service.GetFavoriteList(uid)
	respVideoList := make([]Video, 0)
	copier.Copy(&respVideoList, &videoList)

	resp := FavoriteListResponse{StatusCode: consts.StatusOK, StatusMsg: "返回成功", VideoList: respVideoList}
	ctx.JSON(consts.StatusOK, resp)
}

// FavoriteAction 点赞和取消点赞操作
func FavoriteAction(c context.Context, ctx *app.RequestContext) {
	userId, _ := ctx.Get("user_id")
	videoId, _ := strconv.Atoi(ctx.Query("video_id"))
	actionType, _ := strconv.Atoi(ctx.Query("action_type"))
	res := service.FavoriteAction(userId.(int), videoId, actionType)
	if res {
		ctx.JSON(consts.StatusOK, msg.LikeFavoriteActionSuccess)
	} else {
		ctx.JSON(consts.StatusOK, msg.LikeFavoriteActionFail)
	}
}
