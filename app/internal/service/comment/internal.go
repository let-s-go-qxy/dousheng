package comment

import (
	repository "tiktok/app/internal/model"
	"time"
)

// GetCommentList查选该视频下的所有评论
func GetCommentList(videoId int) (comments []repository.Comment, vidoeCommentCount int) {

	// 调用model层comment的sql查询语句，根据视频id查询对应id的视频评论
	comments = repository.FindCommentByVideo(videoId)
	// 得到视频下的评论数
	vidoeCommentCount = len(comments)
	return
}
func CommentAction(videoId int, actionType int, content string, commentId int, token string) (comment repository.Comment) {
	//var comment repository.Comment
	comment.Content = content
	comment.VideoId = commentId
	comment.VideoId = videoId
	comment.Cancel = actionType
	// TODO 用户赋值
	//comment.User=
	if actionType == 1 {
		// 设定创建时间
		comment.CreateTime = time.Now().Format("01-02")
		repository.CreateComment(comment)
		comment := repository.FindCommentById(commentId)
		return comment
	} else {
		repository.DeleteComment(commentId)
		comment := repository.FindCommentById(commentId)
		return comment
	}
}
