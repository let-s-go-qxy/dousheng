package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
	"tiktok/app/internal/service/ralation"
	userService "tiktok/app/internal/service/user"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// GetFollowAction 关注操作
func GetFollowAction(c context.Context, ctx *app.RequestContext) {
	userId, _ := strconv.Atoi(ctx.Query("user_id"))
	followId, _ := strconv.Atoi(ctx.Query("follow_id"))
	Cancel, _ := strconv.Atoi(ctx.Query("cancel"))
	err := ralation.GetFollowAction(userId, followId, Cancel)
	if err != nil {
		ctx.JSON(consts.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "",
		})
		ctx.JSON(consts.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  "关注成功",
		})

	}

}

// GetFollowerList 获取关注列表
func GetFollowerList(c context.Context, ctx *app.RequestContext) {
	userId, _ := strconv.Atoi(ctx.Query("user_id"))
	ids, err := ralation.GetFollowerList(userId)
	followUsers := make([]User, 0)
	for _, id := range ids {
		user := &User{
			Id: id,
		}
		myId, _ := ctx.Get("user_id")
		user.Id, user.FollowCount, user.FollowerCount, user.Name, user.IsFollow, err = userService.UserInfo(myId.(int), user.Id)
		if err != nil {
			break
		}
		followUsers = append(followUsers, *user)
	}
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
		UserList: followUsers,
	})
}

// GetFollowList 获取粉丝列表
func GetFollowList(c context.Context, ctx *app.RequestContext) {
	userId, _ := strconv.Atoi(ctx.Query("user_id"))
	followerId, err := ralation.GetFollowList(userId)
	followerUsers := make([]User, 0)
	for _, id := range followerId {
		user := &User{
			Id: id,
		}
		myId, _ := ctx.Get("user_id")
		user.Id, user.FollowCount, user.FollowerCount, user.Name, user.IsFollow, err = userService.UserInfo(myId.(int), user.Id)
		if err != nil {
			break
		}
		followerUsers = append(followerUsers, *user)
	}
	if err != nil {
		ctx.JSON(consts.StatusOK, UserListResponse{
			Response: Response{
				statuscode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	ctx.JSON(consts.StatusOK, UserListResponse{
		Response: Response{
			statuscode: 0,
			StatusMsg:  "ok",
		},
		UserList: followerUsers,
	})
}

// GetFriendList 获取朋友列表
func GetFriendList(c context.Context, ctx *app.RequestContext) {
	userId, _ := strconv.Atoi(ctx.Query("user_id"))
	followId, _ := strconv.Atoi(ctx.Query("follow_id"))
	friendId, err := ralation.GetFriendList(userId, followId)
	friendUsers := make([]User, 0)
	for _, id := range friendId {
		user := &User{
			Id: id,
		}
		myId, _ := ctx.Get("user_id")
		user.Id, user.FollowCount, user.FollowerCount, user.Name, user.IsFollow, err = userService.UserInfo(myId.(int), user.Id)
		if err != nil {
			break
		}
		friendUsers = append(friendUsers, *user)
	}
	if err != nil {
		ctx.JSON(consts.StatusOK, UserListResponse{
			Response: Response{
				statuscode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	ctx.JSON(consts.StatusOK, UserListResponse{
		Response: Response{
			statuscode: 0,
			StatusMsg:  "ok",
		},
		UserList: friendUsers,
	})
}
