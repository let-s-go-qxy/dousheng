package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jinzhu/copier"
	"strconv"
	"tiktok/app/global"
	sc "tiktok/app/internal/service/comment"
)

// GetCommentList 获取视频id的评论，以评论时间排序
func GetCommentList(c context.Context, ctx *app.RequestContext) {
	// 获取请求参数
	//token := ctx.Query("token")
	videoId := ctx.Query("video_id")
	vid, err := strconv.Atoi(videoId)
	if err != nil {
		global.Logger.Error("视频ID错误")
	}
	// 进行用户鉴权token服务？权限不够返回

	// 1、通过video表查询对应主键服务；2、根据video主键作为查询条件，查询相应评论
	comments, vidoeCommentCount := sc.GetCommentList(vid)
	print(vidoeCommentCount)
	respCommentList := make([]Comment, 0)
	for _, comment := range comments {
		respComment := Comment{}
		copier.Copy(&respComment, &comment)
		respCommentList = append(respCommentList, respComment)
	}
	resp := CommentListResponse{StatusCode: consts.StatusOK, StatusMsg: "返回成功", CommentList: respCommentList}
	ctx.JSON(consts.StatusOK, resp)
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

	comment := sc.CommentAction(videoId, actionType, commentText, commentId, token)

	respComment := Comment{}
	respComment.Id = comment.Id
	//respComment.User = comment.User
	respComment.Content = comment.Content
	respComment.CreateDate = comment.CreateTime
	resp := CommentResponse{StatusCode: consts.StatusOK, StatusMsg: "返回成功", Comment: respComment}
	ctx.JSON(consts.StatusOK, resp)
}
