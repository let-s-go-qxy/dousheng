package api

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strconv"
	"tiktok/app/internal/model"
	"tiktok/app/internal/service/video"
	"tiktok/utils/common"
	"tiktok/utils/msg"
	"time"
)

// GetFeedList 获取Feed列表
func GetFeedList(c context.Context, ctx *app.RequestContext) {
	lastTime, _ := strconv.ParseInt(ctx.Query("last_time"), 10, 32)
	userIDInterface, success := ctx.Get("user_id")
	var userID int32
	if success {
		userID = int32(userIDInterface.(int))
	} // 若不存在，userID默认为0

	if lastTime == 0 {
		lastTime = time.Now().Unix()
	}
	// 需要获取NextTime、VideoList
	nextTime, videoInfo, state := video.GetVideoFeed(lastTime, userID)

	if state == 0 {
		ctx.JSON(http.StatusOK, &model.GetVideoResponse{
			Response: common.Response{
				StatusCode: -1,
				StatusMsg:  msg.HasNoVideoMsg,
			}, NextTime: lastTime,
		})
	} else if state == 1 {
		fmt.Println("----->>", videoInfo[0].Author)
		ctx.JSON(http.StatusOK, &model.GetVideoResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  msg.GetVideoInfoSuccessMsg,
			}, NextTime: nextTime,
			VideoList: videoInfo,
		})
	}
}
