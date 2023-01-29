package model

type Comment struct {
	Id         int    `gorm:"primaryKey" json:"id"`
	User       User   `json:"user"`
	Content    string `json:"content"`
	CreateTime string `json:"create_date"`
	cancel     int
}
