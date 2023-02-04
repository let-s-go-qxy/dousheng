package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
	g "tiktok/app/global"
	"tiktok/app/internal/model"
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
	toid := ctx.Query("to_user_id")
	fromid := ctx.Query("user_id")

	toId, err := strconv.Atoi(toid)
	if err != nil {
		g.Logger.Error("获取对方ID错误")
	}
	fromId, err := strconv.Atoi(fromid)
	if err != nil {
		g.Logger.Error("获取自己ID错误")
	}

	//token鉴权

	//token进行用户鉴权
	ctx.Query("token")

	fromPrimaryKey := g.MysqlDB.First(fromId)
	toPrimaryKey := g.MysqlDB.First(toId)

	var tomessage []model.Message   //发送的消息
	var resmessage []model.Message  //接收的消息
	var messagelist []model.Message //合并之后的消息

	//根据fromid和toid查询对应的消息
	g.MysqlDB.Where("from_user_id = ?", fromPrimaryKey).Find(&tomessage)
	g.MysqlDB.Where("to_user_id = ?", toPrimaryKey).Find(&resmessage)

	//把发送的消息和接收的消息进行合并到messagelist
	i := 0
	j := 0
	for i < len(tomessage) && j < len(resmessage) {
		if tomessage[i].CreateTime > resmessage[j].CreateTime {
			messagelist = append(messagelist, tomessage[i])
			i++
		} else {
			messagelist = append(messagelist, resmessage[j])
			j++
		}
	}
	if i == len(tomessage) {
		for j < len(resmessage) {
			messagelist = append(messagelist, resmessage[j])
			j++
		}
	}
	if j == len(resmessage) {
		for j < len(tomessage) {
			messagelist = append(messagelist, tomessage[j])
			i++
		}
	}

	resp := MessageListResponse{Response: Response{
		StatusCode: g.StatusCodeOk,
		StatusMsg:  "获取消息列表成功!!"},
		MessageList: respMessageList}
  ctx.JSON(consts.statusOk, resp)
}

func GetMessageAction(c context.Context, ctx *app.RequestContext) {
	userIDInterface, success := ctx.Get("user_id")
	var userID int
	if success {
		userID = int(userIDInterface.(int))
	} // 若不存在，userID默认为0

	toId, _ := strconv.Atoi(ctx.Query("to_user_id"))
	content := ctx.Query("content")
	actionType, _ := strconv.Atoi(ctx.Query("action_type"))

	err := m.MessgaeAction(userID, toId, content, actionType)
	if err != nil {
		ctx.JSON(consts.StatusOK, Response{
			StatusCode: g.StatusCodeFail,
			StatusMsg:  err.Error(),
		})
	}

	resp := Response{StatusCode: g.StatusCodeOk, StatusMsg: "发送消息成功"}
	ctx.JSON(consts.StatusOK, resp)

}
