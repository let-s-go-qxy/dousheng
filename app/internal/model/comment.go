package model

import (
	g "tiktok/app/global"
)

type Comment struct {
	Id         int    `gorm:"primaryKey" json:"id"`
	UserId     int    `json:"user_id"`
	Content    string `json:"comment_text" gorm:"column:comment_text"`
	CreateTime string `json:"create_time" gorm:"column:create_time"`
	VideoId    int    `json:"video_id"`
	Cancel     int    `gorm:"default:1"`
}

// 根据视频id查询全部评论，并且仅返回cancel等于1的
func FindCommentByVideo(id int) []Comment {
	comments := make([]Comment, 0)
	g.MysqlDB.Where("video_id = ? AND cancel = 1 ", id).Order("create_time desc").Find(&comments)
	return comments
}

// 根据评论的id查询返回对应评论
func FindCommentById(id int) (comment Comment) {
	g.MysqlDB.Where("id = ?", id).Find(&comment)
	return
}

// comment表内插入对应评论
func CreateComment(comment *Comment) (returnComment *Comment, err error) {
	db := g.MysqlDB.Create(comment)
	model := db.Statement.Model
	returnComment = model.(*Comment)
	return
}

// comment表内删除对应评论,输入参数为comment整体
func DeleteComment(comment *Comment) (err error) {
	err = g.MysqlDB.Model(&Comment{}).Where("id =?", comment.Id).Update("cancel", "0").Error
	return
}
