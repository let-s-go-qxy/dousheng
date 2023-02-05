package model

import (
	//"fmt"
	g "tiktok/app/global"
)

type Message struct {
	Id         int    `json:"id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
}

type MessageSendEvent struct {
	UserId     int    `json:"user_id,"`
	ToUserId   int    `json:"to_user_id"`
	MsgContent string `json:"msg_content"`
}

type MessagePushEvent struct {
	FromUserId int    `json:"user_id"`
	MsgContent string `json:"msg_content"`
}

type RespMessage struct {
	Id         int    `json:"id"`
	ToId       int    `json:"to_id"`
	FromId     int    `json:"from_id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
}

func GetMessageList(toUserId int, fromUserId int) (respMessageList []RespMessage) {

	g.MysqlDB.Table("messages").
		Where("from_id = ? and to_id = ?", fromUserId, toUserId).
		Scan(&respMessageList)
	return
}

/*func GetFromId(toUserId int) (fromUserId int) {
	g.MysqlDB.Table("messages").Select("from_id").
		Where("to_id = ?", toUserId).
		Find(&fromUserId)
	fmt.Printf("%d", fromUserId)
	return
}
*/

// 创建消息
func CreateMessage(message *RespMessage) (err error) {
	err = g.MysqlDB.Table("messages").Create(message).Error
	return
}
