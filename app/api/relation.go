package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	g "tiktok/app/global"
	"tiktok/app/internal/model"
)

// GetFollowerList 获取关注列表
func GetFollowerList(c context.Context, ctx *app.RequestContext) {
	var users []model.User
	g.MysqlDB.Find(&users)
	ctx.JSON(consts.StatusOK, utils.H{"message": "pong", "data": users})
}
