package model

type Message struct {
	Id         int    `gorm:"primaryKey" json:"id"`
	FromId     int    `json:"from_id"`
	ToId       int    `json:"to_id"`
	Content    string `json:"content" gorm:"column:content"`
	CreateTime string `json:"create_time" gorm:"column:create_time"`
}

// CreateMessage
func CreateMessage(message *RespMessage) (err error) {
	err = g.MysqlDB.Table("messages").Create(message).Error
	return
}
