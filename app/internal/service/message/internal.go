package message

import (
	"tiktok/app/internal/model"

	"github.com/jinzhu/copier"
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
