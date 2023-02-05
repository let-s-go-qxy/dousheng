package ralation

import (
	"tiktok/app/internal/model"
)

func GetFollowAction(userId, followId, Cancel int) (id []int) {
	id = model.GetFollowAll(userId, followId, Cancel)
	return
}

func GetFollowerList(userId int) (ids []int, err error) {
	ids = model.GetFollowsByUserId(userId)
	// userId -> 去follow表查所有记录userId -> 去User根据所有的followId对应的用户信息
	return
}

func GetFollowList(userId int) (followerId []int, err error) {
	followerId = model.GetFollowUserId(userId)
	// userId -> 去follow表查所有记录userId -> 去User根据所有的followId对应的用户信息
	return
}
func GetFriendList(userId int, followId int) (friendId []int, err error) {
	friendId = model.GetFriendUseId(userId, followId)
	// userId,followId -> 去follow表查所有记录userId -> 去User根据所有的followId对应的用户信息
	return
}
