package model

type Comment struct {
	Id         int    `gorm:"primaryKey" json:"id"`
	FromId     int    `json:"from_id"`
	ToId       int    `json:"to_id"`
	Content    string `json:"content" gorm:"column:content"`
	CreateTime string `json:"create_time" gorm:"column:create_time"`
}
