package middleware

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"tiktok/app/api"
	"tiktok/app/internal/service"
	"time"
)

func Jwt() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		fmt.Println("jwt")
		token := ctx.Query("token")
		if token2 := ctx.PostForm("token"); token2 != "" {
			token = token2
		}
		if token == "" {
			ctx.AbortWithStatusJSON(consts.StatusOK, api.Response{
				StatusCode: 1, StatusMsg: "User doesn't exist",
			})
			return
		}
		claims, err := service.ParseToken(token)
		if err != nil || claims.Expire < int(time.Now().Unix()) {
			// TODO 这里应该是返回401，但是Demo中是这么写的，所以为了适配app故此
			ctx.AbortWithStatusJSON(consts.StatusOK, api.Response{
				StatusCode: 1, StatusMsg: "User doesn't exist",
			})
			return
		}
		ctx.Set("user_id", claims.Id)
		//if id := ctx.Query("user_id"); id != "" {
		//	if id != strconv.Itoa(claims.Id) {
		//		ctx.AbortWithStatusJSON(consts.StatusOK, api.Response{
		//			StatusCode: 1, StatusMsg: "User doesn't exist",
		//		})
		//		return
		//	}
		//}
		ctx.Next(c)
	}
}
