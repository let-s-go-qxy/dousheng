package api

import (
	"context"
	"strconv"
	g "tiktok/app/global"
	"tiktok/app/internal/model"

	"github.com/cloudwego/hertz/pkg/app"
)

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

	var tomessage []model.MessageSendEvent  //发送的消息
	var resmessage []model.MessagePushEvent //接收的消息
	var messagelist []model.Message         //合并之后的消息

	//根据fromid和toid查询对应的消息
	g.MysqlDB.Where("from_user_id = ?", fromPrimaryKey).Find(&tomessage)
	g.MysqlDB.Where("to_user_id = ?", toPrimaryKey).Find(&resmessage)
	g.MysqlDB.Where("to_user_id = ?", toPrimaryKey).Find(&messagelist)

	//把发送的消息和接收的消息进行合并到messagelist

}
func GetMessageAction(c context.Context, ctx *app.RequestContext) {

}
