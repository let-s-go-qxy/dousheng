package ralation

import (
	"tiktok/app/internal/model"
	"tiktok/app/internal/service/message"
	userSerive "tiktok/app/internal/service/user"
)

type User struct {
	Id            int
	Name          string
	FollowCount   int
	FollowerCount int
	IsFollow      bool
}

type UserAndMsg struct {
	User
	Message string
	MsgType int
}

func RelationAction(myId, toUserId, actionType int) error {
	return model.CreateOrUpdateFollow(myId, toUserId, actionType)
}

// GetFollowList 获取关注者列表
func GetFollowList(userId, myId int) (followUsers []User, err error) {
	ids := model.GetFollowsByUserId(userId)
	for _, id := range ids {
		user := new(User)
		user.Id, user.FollowCount, user.FollowerCount, user.Name, user.IsFollow, err = userSerive.UserInfo(myId, id)
		if err != nil {
			break
		}
		followUsers = append(followUsers, *user)
	}
	// userId -> 去follow表查所有记录userId -> 去User根据所有的followId对应的用户信息
	return
}

// GetFollowerList 获取粉丝列表
func GetFollowerList(userId, myId int) (followerUsers []User, err error) {
	ids := model.GetFollowersByUserId(userId)
	for _, id := range ids {
		user := new(User)
		user.Id, user.FollowCount, user.FollowerCount, user.Name, user.IsFollow, err = userSerive.UserInfo(myId, id)
		if err != nil {
			break
		}
		followerUsers = append(followerUsers, *user)
	}
	// userId -> 去follow表查所有记录userId -> 去User根据所有的followId对应的用户信息
	return
}

// GetFriendList 获取好友列表，带上最后一次聊天记录
func GetFriendList(myId int) (friends []UserAndMsg, err error) {
	ids := model.GetFriendsByUserId(myId)
	for _, id := range ids {
		user := new(UserAndMsg)
		user.Id, user.FollowCount, user.FollowerCount, user.Name, user.IsFollow, err = userSerive.UserInfo(myId, id)
		user.Message, user.MsgType = message.GetMsgLatest(id, myId)
		if err != nil {
			break
		}
		friends = append(friends, *user)
	}
	// userId -> 去follow表查所有记录userId -> 去User根据所有的followId对应的用户信息
	return
}
