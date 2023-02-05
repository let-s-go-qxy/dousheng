package message

import "tiktok/app/internal/model"

// GetMsgLatest 获取最新的聊天记录  msgType 0为接收的信息，1为发送的信息
func GetMsgLatest(userId, myId int) (msg string, msgType int) {
	msg, msgType = model.GetMsgLatest(userId, myId)
	return
}
