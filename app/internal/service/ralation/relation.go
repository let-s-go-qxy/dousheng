package ralation

import (
	"strconv"
	g "tiktok/app/global"
	"tiktok/app/internal/model"
	"tiktok/app/internal/service/message"
	userSerive "tiktok/app/internal/service/user"
	"tiktok/utils/sort"
	"time"
)

type User struct {
	Id            int
	Name          string
	FollowCount   int
	FollowerCount int
	IsFollow      bool
	Avatar        string `json:"avatar"`
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
		user.Avatar = userSerive.GetAvatar(user.Id)
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
		user.Avatar = userSerive.GetAvatar(user.Id)
		if err != nil {
			break
		}
		followerUsers = append(followerUsers, *user)
	}
	// userId -> 去follow表查所有记录userId -> 去User根据所有的followId对应的用户信息
	return
}

//获取当前用户的所有朋友的ID列表
func GetFriendsIdList(myId int) (friendsIdList []int, err error) {
	friends, err := GetFriendList(myId)
	if err != nil {
		g.Logger.Infof("GetFriendIdList函数获取朋友列表时出错了！")
	}
	friendsIdList = make([]int, 0)
	for _, friend := range friends {
		friendId := friend.Id
		friendsIdList = append(friendsIdList, friendId)
	}
	return
}

//得到当前用户所有的聊天记录列表  并按MessageId排序返回
func GetAllFriendsMessageList(myId int) (respMessageList []model.RespMessage, err error) {
	friendsId, _ := GetFriendsIdList(myId)
	allMessageIdList := make([]int, 0)
	messageIdList := make([]int, 0)
	for _, friendId := range friendsId {
		messageSendIdList := model.GetMessageIdList(myId, friendId)
		messageReceiveIdList := model.GetMessageIdList(friendId, myId)
		messageIdList = append(messageSendIdList, messageReceiveIdList...)
		allMessageIdList = append(allMessageIdList, messageIdList...)
	}

	allMessageIdList = sort.QuickSort(allMessageIdList)

	allMessageList := model.GetMessageList(allMessageIdList)
	for i, message := range allMessageList {
		t := time.Time{}
		//fmt.Println("message.CreateTime:---->>" + message.CreateTime)
		t, _ = time.ParseInLocation("2006-01-02T15:04:05Z07:00", message.CreateTime, time.Local)
		//fmt.Println(t)
		allMessageList[i].CreateTime = strconv.Itoa(int(t.Unix()))
		//fmt.Println("allMessageList[i].CreateTime:------->" + allMessageList[i].CreateTime)
	}
	return allMessageList, err
}

// GetFriendList 获取好友列表，带上最后一次聊天记录
func GetFriendList(myId int) (friends []UserAndMsg, err error) {
	ids := model.GetFriendsByUserId(myId)
	for _, id := range ids {
		user := new(UserAndMsg)
		user.Id, user.FollowCount, user.FollowerCount, user.Name, user.IsFollow, err = userSerive.UserInfo(myId, id)
		user.Avatar = userSerive.GetAvatar(user.Id)
		user.Message, user.MsgType = message.GetMsgLatest(id, myId)
		if err != nil {
			break
		}
		friends = append(friends, *user)
	}
	// userId -> 去follow表查所有记录userId -> 去User根据所有的followId对应的用户信息
	return
}
