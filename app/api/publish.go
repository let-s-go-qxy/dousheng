package api

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/json"
	"strconv"
	"tiktok/app/internal/service/video"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jinzhu/copier"
	g "tiktok/app/global"
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
	userId, err := strconv.Atoi(ctx.Query("user_id"))
	if err != nil {
		g.Logger.Error("用户ID错误")
	}

	videoList, _ := video.GetPublicList(userId)

	respVideoList := make([]Video, 0)

	copier.Copy(&respVideoList, &videoList)
	resp := VideoListResponse{Response: Response{
		StatusCode: g.StatusCodeOk, StatusMsg: "成功!!"},
		VideoList: respVideoList}
	marshal, _ := json.Marshal(respVideoList)
	fmt.Println(string(marshal))
	ctx.JSON(consts.StatusOK, resp)
}
