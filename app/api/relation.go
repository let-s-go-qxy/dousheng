package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jinzhu/copier"
	"strconv"
	"tiktok/app/internal/service/ralation"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

type UserAndMsg struct {
	User
	Message string `json:"message"`
	MsgType int    `json:"msg_type"`
}

type UserAndMsgListResponse struct {
	Response
	UserList []UserAndMsg `json:"user_list"`
}

// RelationAction 关注
func RelationAction(c context.Context, ctx *app.RequestContext) {
	toUserId, _ := strconv.Atoi(ctx.Query("to_user_id"))
	actionType, _ := strconv.Atoi(ctx.Query("action_type"))
	myId, _ := ctx.Get("user_id")
	err := ralation.RelationAction(myId.(int), toUserId, actionType)
	if err != nil {
		ctx.JSON(consts.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	ctx.JSON(consts.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "ok",
	})
}

// GetFollowerList 获取粉丝列表
func GetFollowerList(c context.Context, ctx *app.RequestContext) {
	userId, _ := strconv.Atoi(ctx.Query("user_id"))
	myId, _ := ctx.Get("user_id")
	followerUsers, err := ralation.GetFollowerList(userId, myId.(int))
	if err != nil {
		ctx.JSON(consts.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	followerUsersResp := make([]User, 0)
	err = copier.Copy(&followerUsersResp, &followerUsers)
	if err != nil {
		ctx.JSON(consts.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	ctx.JSON(consts.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "ok",
		},
		UserList: followerUsersResp,
	})
}

// GetFollowList 获取关注者列表
func GetFollowList(c context.Context, ctx *app.RequestContext) {
	userId, _ := strconv.Atoi(ctx.Query("user_id"))
	myId, _ := ctx.Get("user_id")
	followUsers, err := ralation.GetFollowList(userId, myId.(int))
	if err != nil {
		ctx.JSON(consts.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	followUsersResp := make([]User, 0)
	err = copier.Copy(&followUsersResp, &followUsers)
	if err != nil {
		ctx.JSON(consts.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	ctx.JSON(consts.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "ok",
		},
		UserList: followUsersResp,
	})
}

// GetFriendList 获取好友列表 同时获取最新的聊天记录
func GetFriendList(c context.Context, ctx *app.RequestContext) {
	myId, _ := ctx.Get("user_id")
	friendUsers, err := ralation.GetFriendList(myId.(int))
	if err != nil {
		ctx.JSON(consts.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	friendUsersResp := make([]UserAndMsg, 0)
	err = copier.Copy(&friendUsersResp, &friendUsers)
	if err != nil {
		ctx.JSON(consts.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	ctx.JSON(consts.StatusOK, UserAndMsgListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "ok",
		},
		UserList: friendUsersResp,
	})
}
