package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	g "tiktok/app/global"
	"tiktok/app/internal/model"
)

// GetCommentList 获取视频id的评论，以评论时间排序
func GetCommentList(c context.Context, ctx *app.RequestContext) {
	// 获取请求参数
	//token := ctx.Query("token")
	videoId := ctx.Query("video_id")
	// 进行用户鉴权token服务？权限不够返回

	// 通过video表查询对应主键服务
	videoPrimaryKey := g.MysqlDB.First(videoId)

	// 根据video主键作为查询条件，查询相应评论
	var comments []model.Comment

	// 根据视频id查询对应id的视频评论
	g.MysqlDB.Where("video_id = ?", videoPrimaryKey).Find(&comments)

	ctx.JSON(consts.StatusOK, utils.H{"message": "成功", "comment_list": comments})
}

func GetCommentAction(c context.Context, ctx *app.RequestContext) {
	// 获取请求参数
	//videoId := ctx.Query("video_id")
	//token := ctx.Query("token")
	//actionType := ctx.Query("action_type")
	//commentText := ctx.Query("comment_text")
	//commentId := ctx.Query("comment_id")

	ctx.JSON(consts.StatusOK, utils.H{"message": "成功", "comment_list": "ok"})
}
