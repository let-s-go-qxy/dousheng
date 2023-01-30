package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
	g "tiktok/app/global"
	repository "tiktok/app/internal/model"
	sc "tiktok/app/internal/service/comment"
)

type CommentResponse struct {
	Response
	Comments []repository.Comment `json:"comment_list"`
}

// GetCommentList 获取视频id的评论，以评论时间排序
func GetCommentList(c context.Context, ctx *app.RequestContext) {
	// 获取请求参数
	//token := ctx.Query("token")
	videoId := ctx.Query("video_id")
	// 进行用户鉴权token服务？权限不够返回

	// 1、通过video表查询对应主键服务；2、根据video主键作为查询条件，查询相应评论
	sc.GetComment(videoId)
	ctx.JSON(consts.StatusOK, CommentResponse{
		Response: Response{
			StatusCode: g.StatusCodeOk,
		},
		Comments: sc.GetComment(videoId),
	})
}

// PostCommentAction对视频下的评论进行发表或者删除
func PostCommentAction(c context.Context, ctx *app.RequestContext) {
	// 获取请求参数
	token := ctx.Query("token")
	videoId, _ := strconv.Atoi(ctx.Query("video_id"))       //》视频肯定存在》根据视频查找对应评论
	actionType, _ := strconv.Atoi(ctx.Query("action_type")) //》视频操作？1》添加insert：2》删除delete
	commentText := ctx.Query("comment_text")
	commentId, _ := strconv.Atoi(ctx.Query("comment_id"))
	// 进行用户鉴权token服务？权限不够返回
	sc.CommentAction(videoId, actionType, commentText, commentId, token)
	ctx.JSON(consts.StatusOK, utils.H{"message": "成功", "comment_list": "ok"})
}
