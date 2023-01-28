package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetFollowerList 获取关注列表
func GetFollowerList(c context.Context, ctx *app.RequestContext) {
	ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
}
