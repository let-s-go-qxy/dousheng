package message

import (
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/streadway/amqp"
	"strconv"
	g "tiktok/app/global"
	"tiktok/app/internal/model"
	"tiktok/utils/mq"
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
		//fmt.Println(message.CreateTime)
		t, _ = time.ParseInLocation("2006-01-02T15:04:05Z07:00", message.CreateTime, time.Local)
		//fmt.Println(t)
		respMessageList[i].CreateTime = strconv.Itoa(int(t.Unix()))
	}
	return respMessageList, err

}

func GetRabbitMQMessageList(userId int) (respMessageList []model.RespMessage, err error) {
	conn, _ := amqp.Dial("amqp://admin:Qd20010701.@10.211.55.4:5672/")
	strUserId := strconv.Itoa(userId)
	ch, _ := conn.Channel()
	argumentsMap := map[string]interface{}{}
	argumentsMap["x-max-length"] = 1
	argumentsMap["x-overflow"] = "drop-head"
	q, _ := ch.QueueDeclare(
		"message_list"+strUserId, // name
		true,                     // durable
		false,                    // delete when unused
		false,                    // exclusive
		false,                    // no-wait
		argumentsMap,             // arguments
	)
	msgs, _ := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	allMessageList := []model.RespMessage{}
	messageList := []model.RespMessage{}
	//message := model.RespMessage{}

	go func() {
		for d := range msgs {
			//json.Unmarshal(d.Body, &message)
			json.Unmarshal(d.Body, &messageList)
			//messageList = append(messageList, message)
			allMessageList = append(allMessageList, messageList...)
		}
	}()
	err = ch.Close()
	if err != nil {
		g.Logger.Infof("ch.Close()时发生了错误！")
	}
	err = conn.Close()
	if err != nil {
		g.Logger.Infof("conn.Close()时发生了错误！")
	}
	return allMessageList, err
}

func GetRabbitMQMessageCurrent(userId int) (respMessageList []model.RespMessage, err error) {
	strUserId := strconv.Itoa(userId)
	conn, _ := amqp.Dial("amqp://admin:Qd20010701.@10.211.55.4:5672/")

	ch, _ := conn.Channel()

	argumentsMap := map[string]interface{}{}
	argumentsMap["x-max-length"] = 1
	argumentsMap["x-overflow"] = "drop-head"
	q, _ := ch.QueueDeclare(
		"message_current"+strUserId, // name
		true,                        // durable
		false,                       // delete when unused
		false,                       // exclusive
		false,                       // no-wait
		argumentsMap,                // arguments
	)
	msgs, _ := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	messageList := []model.RespMessage{}
	message := model.RespMessage{}

	go func() {
		for d := range msgs {
			json.Unmarshal(d.Body, &message)
			messageList = append(messageList, message)
		}
	}()
	err = ch.Close()
	if err != nil {
		g.Logger.Infof("ch.Close()时发生了错误！")
	}
	err = conn.Close()
	if err != nil {
		g.Logger.Infof("conn.Close()时发生了错误！")
	}
	return messageList, err
}

func MessageAction(fromId int, toId int, content string, actionType int) (err error) {

	msg := model.RespMessage{
		ToId:       toId,
		FromId:     fromId,
		Content:    content,
		CreateTime: time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}
	if actionType == g.MessageSendEvent {
		err = model.CreateMessage(&msg)
		if err != nil {
			err = errors.New("发送消息失败: " + err.Error())
		}
	}
	fmt.Println(msg.CreateTime)

	t := time.Time{}
	//fmt.Println(message.CreateTime)
	t, _ = time.ParseInLocation("2006-01-02T15:04:05Z07:00", msg.CreateTime, time.Local)
	msg.CreateTime = strconv.Itoa(int(t.Unix()))

	JsonMsg, err := json.Marshal(msg)
	strJsonMsg := string(JsonMsg)

	//将消息写到userId对应的消息队列中去
	//mq.PublishMessageCurrentToMQ(strJsonMsg, fromId)
	//将消息写到ToId对应的消息队列中去
	mq.PublishMessageCurrentToMQ(strJsonMsg, toId)

	return
}
