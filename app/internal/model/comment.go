package model

import (
	g "tiktok/app/global"
)

type Comment struct {
	Id         int    `gorm:"primaryKey" json:"id"`
	User       User   `json:"user"`
	Content    string `json:"content"`
	CreateTime string `json:"create_date"`
	VideoId    int    `json:"video_id"`
	Cancel     int    `gorm:"default:1"`
}

func FindCommentByVideo(id int) []Comment {
	comments := make([]Comment, 0)
	g.MysqlDB.Where("video_id = ?", id).Find(&comments)
	return comments
}

func FindCommentById(id int) (comment Comment) {
	g.MysqlDB.Where("id = ?", id).Find(&comment)
	return comment
}

// comment表内插入对应评论
func CreateComment(comment Comment) {
	g.MysqlDB.Create(comment)
}

// comment表内删除对应评论,输入参数为评论的id
func DeleteComment(commentId int) {
	g.MysqlDB.Model(&Comment{Id: commentId}).Update("cancel", "0")
}
