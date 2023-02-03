package api

import (
	"context"
	"fmt"
	"strconv"
	g "tiktok/app/global"
	m "tiktok/app/internal/service/message"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jinzhu/copier"
)

type MergeMessage struct {
	Id         int    `json:"id"`
	ToId       int    `json:"to_id"`
	FromId     int    `json:"from_id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
}
type MessageListResponse struct {
	Response
	MessageList []MergeMessage `json:"message_list"`
}

func GetMessageList(c context.Context, ctx *app.RequestContext) {
	//token鉴权

	//获取to_user_id和user_id

	toId, err := strconv.Atoi(ctx.Query("to_user_id"))
	if err != nil {
		g.Logger.Error("获取对方ID错误")
	}

	fromId := m.GetFromId(toId)

	messageList, _ := m.GetMessageList(toId, fromId)

	respMessageList := make([]MergeMessage, 0)

	copier.Copy(&respMessageList, &messageList)

	resp := MessageListResponse{Response: Response{
		StatusCode: 0,
		StatusMsg:  "成功!!"},
		MessageList: respMessageList}

	marshal, _ := json.Marshal(respMessageList)
	fmt.Println(string(marshal))
	ctx.JSON(consts.StatusOK, resp)
}

func GetMessageAction(c context.Context, ctx *app.RequestContext) {

}
