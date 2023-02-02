package model

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
