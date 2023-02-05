package model

import (
	"fmt"
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

func GetFromId(toUserId int) (fromUserId int) {
	g.MysqlDB.Table("messages").Select("from_id").
		Where("to_id = ?", toUserId).
		Find(&fromUserId)
	fmt.Printf("%d", fromUserId)
	return
}

// 创建消息
func CreateMessage(message *RespMessage) (err error) {
	err = g.MysqlDB.Table("messages").Create(message).Error
	return
}

// GetMsgLatest 获取最新的聊天记录  msgType 0为接收的信息，1为发送的信息
func GetMsgLatest(userId, myId int) (msg string, msgType int) {
	msgDao := new(RespMessage)
	g.MysqlDB.Table("messages").Where("to_id = ? AND from_id = ?", userId, myId).
		Or("to_id = ? AND from_id = ?", myId, userId).Order("create_time desc").First(msgDao)
	if msgDao.ToId == userId {
		msgType = 1
	}
	msg = msgDao.Content
	return
}
