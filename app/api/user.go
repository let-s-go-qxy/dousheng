package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
	g "tiktok/app/global"
	userService "tiktok/app/internal/service/user"
)

type UserLoginResponse struct {
	Response
	UserId int    `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

// UserInfo 获取用户详情
func UserInfo(c context.Context, ctx *app.RequestContext) {
	user := new(User)
	user.Id, _ = strconv.Atoi(ctx.Query("user_id"))
	var err error
	myID, _ := ctx.Get("user_id")
	_, user.FollowCount, user.FollowerCount, user.Name, user.IsFollow, err = userService.UserInfo(myID.(int), user.Id)
	user.Avatar = userService.GetAvatar(user.Id)
	if err != nil {
		ctx.JSON(consts.StatusOK, Response{
			StatusCode: g.StatusCodeFail,
			StatusMsg:  err.Error(),
		})
	}
	ctx.JSON(consts.StatusOK, UserResponse{
		Response: Response{
			StatusCode: g.StatusCodeOk,
		},
		User: *user,
	})
}

// UserRegister 用户注册
func UserRegister(c context.Context, ctx *app.RequestContext) {
	name := ctx.Query("username")
	pw := ctx.Query("password")
	userId, token, err := userService.UserRegister(name, pw)
	if err != nil {
		ctx.JSON(consts.StatusOK, Response{
			StatusCode: g.StatusCodeFail,
			StatusMsg:  err.Error(),
		})
		return
	}
	ctx.JSON(consts.StatusOK, UserLoginResponse{
		Response: Response{
			StatusCode: g.StatusCodeOk,
		},
		UserId: userId,
		Token:  token,
	})
}

// UserLogin 用户登录
func UserLogin(c context.Context, ctx *app.RequestContext) {
	name := ctx.Query("username")
	pw := ctx.Query("password")
	userId, token, err := userService.UserLogin(name, pw)
	if err != nil {
		ctx.JSON(consts.StatusOK, Response{
			StatusCode: g.StatusCodeFail,
			StatusMsg:  err.Error(),
		})
		return
	}
	ctx.JSON(consts.StatusOK, UserLoginResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserId: userId,
		Token:  token,
	})
}
