package message

import (
	"errors"
	"fmt"
	"strconv"
	g "tiktok/app/global"
	"tiktok/app/internal/model"
	"tiktok/utils/sort"
	"time"

	"github.com/jinzhu/copier"
)

func GetMessageList(toUserId int, fromUserId int) (respMessageList []model.RespMessage, err error) {

	var messageList []model.RespMessage

	messageSendIdList := model.GetMessageIdList(toUserId, fromUserId)
	messageReceiveIdList := model.GetMessageIdList(fromUserId, toUserId)
	messageIdList := append(messageSendIdList, messageReceiveIdList...)

	messageIdList = sort.QuickSort(messageIdList)

	messageList = model.GetMessageList(messageIdList)

	for _, message := range messageList {
		respMessage := model.RespMessage{}
		copier.Copy(&respMessage, &message)
		respMessageList = append(respMessageList, respMessage)
	}

	for i, message := range respMessageList {
		t := time.Time{}
		fmt.Println(message.CreateTime)
		t, _ = time.ParseInLocation("2006-01-02T15:04:05Z07:00", message.CreateTime, time.Local)
		fmt.Println(t)
		respMessageList[i].CreateTime = strconv.Itoa(int(t.Unix()))
	}
	return

}

func MessageAction(fromId int, toId int, content string, actionType int) (err error) {

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
