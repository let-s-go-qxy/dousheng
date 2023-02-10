package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jinzhu/copier"
	"strconv"
	g "tiktok/app/global"
	"tiktok/app/internal/model"
	m "tiktok/app/internal/service/message"
)

type MergeMessage struct {
	Id         int    `json:"id"`
	ToId       int    `json:"to_user_id"`
	FromId     int    `json:"from_user_id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
}
type MessageListResponse struct {
	Response
	MessageList []MergeMessage `json:"message_list"`
}

func GetMessageList(c context.Context, ctx *app.RequestContext) {

	//获取to_user_id和user_id

	toId, err := strconv.Atoi(ctx.Query("to_user_id"))
	if err != nil {
		g.Logger.Error("获取对方ID错误")
	}

	userIdInterface, success := ctx.Get("user_id")
	var fromId int
	if success {
		fromId = int(userIdInterface.(int))
	} // 若不存在，userID默认为0

	recordList, err := m.GetRabbitMQMessageList(fromId)
	currentList, err := m.GetRabbitMQMessageCurrent(fromId)
	allMessageList := append(recordList, currentList...)
	messageList := []model.RespMessage{}
	for _, message := range allMessageList {
		if (message.ToId == toId && message.FromId == fromId) || (message.ToId == fromId && message.FromId == toId) {
			messageList = append(messageList, message)
		}
	}
	if err != nil {
		g.Logger.Infof("GetRabbitMQMessageList时发生了错误！")
	}
	//messageList, _ := m.GetMessageList(toId, fromId)
	respMessageList := make([]MergeMessage, 0)
	copier.Copy(&respMessageList, &messageList)
	resp := MessageListResponse{Response: Response{
		StatusCode: g.StatusCodeOk,
		StatusMsg:  "获取消息列表成功!!"},
		MessageList: respMessageList}

	//marshal, _ := json.Marshal(respMessageList)
	//fmt.Println(string(marshal))
	ctx.JSON(consts.StatusOK, resp)

}

func GetMessageAction(c context.Context, ctx *app.RequestContext) {

	userIDInterface, success := ctx.Get("user_id")
	var fromId int
	if success {
		fromId = int(userIDInterface.(int))
	} // 若不存在，userID默认为0

	toId, _ := strconv.Atoi(ctx.Query("to_user_id"))
	content := ctx.Query("content")
	actionType, _ := strconv.Atoi(ctx.Query("action_type"))

	err := m.MessageAction(fromId, toId, content, actionType)
	if err != nil {
		ctx.JSON(consts.StatusOK, Response{
			StatusCode: g.StatusCodeFail,
			StatusMsg:  err.Error(),
		})
	}

	resp := Response{StatusCode: g.StatusCodeOk, StatusMsg: "发送消息成功"}
	ctx.JSON(consts.StatusOK, resp)

}
