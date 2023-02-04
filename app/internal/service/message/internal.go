package message

import (
	"errors"
	"github.com/jinzhu/copier"
	g "tiktok/app/global"
	"tiktok/app/internal/model"
	"time"
)

func GetMessageList(toUserId int, fromUserId int) (respMessageList []model.RespMessage, err error) {

	var messageList []model.RespMessage
	messageList = model.GetMessageList(toUserId, fromUserId)

	for _, message := range messageList {
		respMessage := model.RespMessage{}
		copier.Copy(&respMessage, &message)
		respMessageList = append(respMessageList, respMessage)
	}
	return

}

func GetFromId(toUserId int) (fromUserId int) {
	fromUserId = model.GetFromId(toUserId)
	return
}

func MessgaeAction(fromId int, toId int, content string, actionType int) (err error) {

	msg := model.RespMessage{
		ToId:       toId,
		FromId:     fromId,
		Content:    content,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}

	if actionType == g.MessageSendEvent {
		err = model.CreateMessage(&msg)
		if err != nil {
			err = errors.New("发送消息失败: " + err.Error())
		}
	}
	return
}
