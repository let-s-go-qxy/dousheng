package model

import g "tiktok/app/global"

type Follow struct {
	Id       int `gorm:"primaryKey" json:"id"`
	UserId   int `json:"user_id"`
	FollowId int `json:"follow_id"`
	cancel   int `json:"cancel"`
}

// Author 用户返回模型
type Author struct {
	Id            int    `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int    `json:"follow_count,omitempty"`
	FollowerCount int    `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow"`
}

// GetFollowerCount 获取粉丝数目
func (follow *Follow) GetFollowerCount(userId int) (followerCount int64) {
	g.MysqlDB.Where("follow_id = ?", userId).Count(&followerCount)
	return
}

// GetFollowCount 获取关注数
func (follow *Follow) GetFollowCount(userId int) (followCount int64) {
	g.MysqlDB.Where("user_id = ?", userId).Count(&followCount)
	return
}

// IsFollowed user是否关注另外一个user
func (follow *Follow) IsFollowed(userId int, followId int) bool {
	var count int64
	g.MysqlDB.Where("user_id = ? and follow_id = ?", userId, followId).Count(&count)
	return count > 0
}
