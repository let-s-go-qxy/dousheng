package api

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// PublishList 发布列表
func PublishList(c context.Context, ctx *app.RequestContext) {

	ctx.JSON(consts.StatusOK, utils.H{})
}
