package model

import (
	"gorm.io/gorm"
	g "tiktok/app/global"
	"time"
)

type Comment struct {
	Id         int    `gorm:"primaryKey" json:"id"`
	User       User   `json:"user"`
	Content    string `json:"content"`
	CreateTime string `json:"create_date"`
	VideoId    int    `json:"video_id"`
	Cancel     int    `gorm:"default:1"`
}

func FindCommentByVideo(id *gorm.DB) []Comment {
	comments := make([]Comment, 0)
	g.MysqlDB.Where("video_id = ?", id).Find(&comments)
	return comments
}

// comment表内插入对应评论
func CreateComment(comment Comment) {
	// 设定创建时间
	comment.CreateTime = time.Now().Format("01-02")
	g.MysqlDB.Create(comment)
}

// comment表内删除对应评论,输入参数为评论的id
func DeleteComment(commentId int) {
	g.MysqlDB.Model(&Comment{Id: commentId}).Update("cancel", "0")
}
