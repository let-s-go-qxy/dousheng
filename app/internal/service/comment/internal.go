package comment

import (
	g "tiktok/app/global"
	repository "tiktok/app/internal/model"
)

func GetComment(videoId string) []repository.Comment {

	var comments []repository.Comment
	// 通过video表查询对应主键服务》这里可以调用video的service方法
	videoPrimaryKey := g.MysqlDB.First(videoId)

	// 调用model层comment的sql查询语句，根据视频id查询对应id的视频评论
	comments = repository.FindCommentByVideo(videoPrimaryKey)
	return comments
}
func CommentAction(videoId int, actionType int, content string, commentId int, token string) {
	var comment repository.Comment
	comment.Content = content
	comment.VideoId = commentId
	comment.VideoId = videoId
	comment.Cancel = actionType
	// TODO 用户赋值
	//comment.User=
	if actionType == 1 {
		repository.CreateComment(comment)
	} else if actionType == 2 {
		repository.DeleteComment(commentId)
	}
}
