package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
	"tiktok/app/internal/service"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// GetFollowerList 获取关注列表
func GetFollowerList(c context.Context, ctx *app.RequestContext) {
	userId, _ := strconv.Atoi(ctx.Query("user_id"))
	ids, err := service.GetFollowerList(userId)
	followUsers := make([]User, 0)
	for _, id := range ids {
		user := &User{
			Id: id,
		}
		myId, _ := ctx.Get("user_id")
		user.Id, user.FollowCount, user.FollowerCount, user.Name, user.IsFollow, err = service.UserInfo(myId.(int), user.Id)
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
