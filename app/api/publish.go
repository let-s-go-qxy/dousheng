package api

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/json"
	"strconv"
	"tiktok/app/internal/service/video"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	//"fmt"
	"github.com/jinzhu/copier"
	//"tiktok/app/internal/model"
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

	videoList, _ := video.GetPublicList(userId)

	respVideoList := make([]Video, 0)
	/*
		for _, video := range videoList {
			respVideo := Video{}
			copier.Copy(&respVideo, &video)
			respVideo.FavoriteCount = 0
			respVideo.CommentCount = 0
			respVideo.IsFavorite = false
			fmt.Printf("%+v", respVideo)
			respVideoList = append(respVideoList, respVideo)
		}
	*/
	copier.Copy(&respVideoList, &videoList)
	resp := VideoListResponse{Response: Response{
		StatusCode: 0, StatusMsg: "成功!!"},
		VideoList: respVideoList}
	marshal, _ := json.Marshal(respVideoList)
	fmt.Println(string(marshal))
	ctx.JSON(consts.StatusOK, resp)
}
