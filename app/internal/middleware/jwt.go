package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
	"tiktok/app/api"
	"tiktok/app/internal/service"
	"time"
)

func Jwt() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := c.Query("token")
		if token2 := c.PostForm("token"); token2 != "" {
			token = token2
		}
		if token == "" {
			c.AbortWithStatusJSON(consts.StatusOK, api.Response{
				StatusCode: 1, StatusMsg: "User doesn't exist",
			})
			return
		}
		claims, err := service.ParseToken(token)
		if err != nil || claims.Expire < int(time.Now().Unix()) {
			// TODO 这里应该是返回401，但是Demo中是这么写的，所以为了适配app故此
			c.AbortWithStatusJSON(consts.StatusOK, api.Response{
				StatusCode: 1, StatusMsg: "User doesn't exist",
			})
			return
		}
		c.Set("user_id", claims.Id)
		if id := c.Query("user_id"); id != "" {
			if id != strconv.Itoa(claims.Id) {
				c.AbortWithStatusJSON(consts.StatusOK, api.Response{
					StatusCode: 1, StatusMsg: "User doesn't exist",
				})
				return
			}
		}
		c.Next(ctx)
	}
}
