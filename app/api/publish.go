package api

import (
	"context"
	"strconv"
	"tiktok/app/internal/service"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jinzhu/copier"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// PublishList 发布列表
func PublishList(c context.Context, ctx *app.RequestContext) {
	//token
	println("进入")
	userId, _ := strconv.Atoi(ctx.Query("user_id"))
	videoList, _ := service.GetPublicList(userId)
	respVideoList := make([]Video, 0)

	for _, video := range videoList {
		respVideo := Video{}
		copier.Copy(&respVideo, &video)
		respVideo.FavoriteCount = 0
		respVideo.CommentCount = 0
		respVideo.IsFavorite = false
		respVideoList = append(respVideoList, respVideo)
	}

	resp := VideoListResponse{Response: Response{StatusCode: 0, StatusMsg: "成功!!"}, VideoList: respVideoList}
	ctx.JSON(consts.StatusOK, resp)
}
