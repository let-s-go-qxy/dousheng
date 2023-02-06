package model

import (
	"gorm.io/gorm"
	g "tiktok/app/global"
)

type Follow struct {
	Id       int `gorm:"primaryKey" json:"id"`
	UserId   int `json:"user_id"`
	FollowId int `json:"follow_id"`
	Cancel   int `json:"cancel"`
}

// Author 用户返回模型
type Author struct {
	Id            int    `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int    `json:"follow_count,omitempty"`
	FollowerCount int    `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow"`
}

// GetFollowsByUserId  查所有的被关注者的id
func GetFollowsByUserId(userId int) (arr []int) {
	follows := new([]Follow)
	g.MysqlDB.Find(follows, "user_id = ? AND cancel = ?", userId, 1)
	for _, follow := range *follows {
		arr = append(arr, follow.FollowId)
	}
	return
}

// GetFollowersByUserId 获取所有粉丝的id
func GetFollowersByUserId(userId int) (arr []int) {
	follows := new([]Follow)
	g.MysqlDB.Find(follows, "follow_id = ? AND cancel = ?", userId, 1)
	for _, follow := range *follows {
		arr = append(arr, follow.FollowId)
	}
	return
}

func GetFriendsByUserId(userId int) (arr []int) {
	follows := new([]Follow)
	g.MysqlDB.Find(follows, "user_id = ? AND cancel = ?", userId, 1)
	for _, follow := range *follows {
		var count int64
		g.MysqlDB.Model(&Follow{}).
			Where("user_id = ? AND cancel = ?", follow.FollowId, 1).
			Where("follow_id = ? AND cancel = ?", userId, 1).Count(&count)
		if count > 0 {
			arr = append(arr, follow.FollowId)
		}
	}
	return
}

// GetFollowCount 获取当前用户的关注人数
func GetFollowCount(userId int) (count int64) {
	g.MysqlDB.Model(&Follow{}).Where("user_id = ? AND cancel = ?", userId, 1).Count(&count)
	return
}

// GetFollowerCount 获取当前用户的粉丝人数
func GetFollowerCount(followId int) (count int64) {
	g.MysqlDB.Model(&Follow{}).Where("follow_id = ? AND cancel = ?", followId, 1).Count(&count)
	return
}

func IsFollow(userId, followId int) bool {
	err := g.MysqlDB.First(&Follow{}, "user_id = ? AND follow_id = ? AND cancel = ?", userId, followId, 1).Error
	return err == nil
}

// CreateOrUpdateFollow 新增或更新记录
func CreateOrUpdateFollow(myId, userId, followType int) error {
	follow := new(Follow)
	// 如果有记录更新记录，没有新增记录
	if err := g.MysqlDB.First(follow, "user_id = ? AND follow_id = ?", myId, userId).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
		follow.UserId = myId
		follow.FollowId = userId
		follow.Cancel = followType
		return g.MysqlDB.Create(follow).Error
	}
	return g.MysqlDB.Model(follow).
		Where("user_id = ? AND follow_id = ?", myId, userId).
		Update("cancel", followType).Error
}
