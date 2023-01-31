package model

import (
	"strconv"
	g "tiktok/app/global"
)

type Like struct {
	Id      int `gorm:"primaryKey" json:"id"`
	UserId  int `json:"user_id"`
	VideoId int `json:"video_id"`
	Cancel  int `json:"cancel"`
}

// GetFavoriteList 用户点赞列表
func GetFavoriteList(userId int) (videoList []Video) {
	g.MysqlDB.Model(&Like{}).
		Joins("left join videos on likes.video_id = videos.id").
		Where("likes.id = ? and cancel = ?", userId, g.Cancel).
		Scan(&videoList)
	return
}

// GetFavoriteVideoIdList 获取用户点赞的视频ID列表
func GetFavoriteVideoIdList(userId int) (videoList []int) {
	println("查询点赞列表")
	g.MysqlDB.Table("likes").Select("video_id").Where("user_id = ? and cancel = ?", userId, g.NoCancel).Find(&videoList)
	return
}

// GetVideoFavoriteCount 获取视频点赞数
func GetVideoFavoriteCount(videoId []int) (videoFavoriteCount map[int]int) {
	videoFavoriteCount = map[int]int{}
	for _, id := range videoId {
		count := g.DbVideoLike.LLen(g.RedisContext, strconv.Itoa(id)).Val()
		videoFavoriteCount[id] = int(count)
	}
	return
}
