package service

import (
	"tiktok/app/internal/model"
)

func GetFollowerList(userId int) (ids []int, err error) {
	ids = model.GetFollowsByUserId(userId)
	// userId -> 去follow表查所有记录userId -> 去User根据所有的followId对应的用户信息
	return
}

func GetFollowList(userId int) (followerId []int, err error) {
	followerId = model.GetFollowerUserId(userId)
	// userId -> 去follow表查所有记录userId -> 去User根据所有的followId对应的用户信息
	return
}
