package model

import (
	"strconv"
	g "tiktok/app/global"
	"tiktok/utils"
	"time"
)

type Like struct {
	Id      int `gorm:"primaryKey" json:"id"`
	UserId  int `json:"user_id"`
	VideoId int `json:"video_id"`
	Cancel  int `json:"cancel"`
}

//// GetFavoriteList 用户点赞列表
//func (like *Like) GetFavoriteList(userId int) (videoList []Video) {
//	g.MysqlDB.Model(&Like{}).
//		Joins("left join videos on likes.video_id = videos.id").
//		Where("likes.id = ? and cancel = ?", userId, g.Cancel).
//		Scan(&videoList)
//	return
//}

// GetFavoriteVideoIdList 获取用户点赞的视频ID列表
func (like *Like) GetFavoriteVideoIdList(userId int) (videoList []int) {
	g.MysqlDB.Table("likes").Select("video_id").Where("user_id = ? and cancel = ?", userId, g.NoCancel).Find(&videoList)
	return
}

func (like *Like) GetUserIdListForVideo(videoId int) (userIdList []int) {
	g.MysqlDB.Table("likes").Select("user_id").Where("video_id = ? and cancel = ?", videoId, g.NoCancel).Find(&userIdList)
	return
}

// GetUserFavoriteVideoList 获取视频对应的点赞列表
func (like *Like) GetUserFavoriteVideoList(userId int) (videoIdList []int) {
	exist := g.DbVideoLike.SIsMember(g.RedisContext, "userSet", userId).Val()
	// 缓存不存在
	if !exist {
		// 1、从like表中查询出喜欢的视频ID列表
		videoIdList = like.GetFavoriteVideoIdList(userId)
		// 2、将喜欢视频ID列表缓存到redis中
		g.DbVideoLike.SAdd(g.RedisContext, "userSet", userId)
		g.DbVideoLike.Set(g.RedisContext, strconv.Itoa(userId), videoIdList, time.Hour*24) //缓存时间设置？
	} else {
		videoIdListTmp := g.DbVideoLike.LRange(g.RedisContext, string(userId), 0, -1).Val()
		videoIdList = utils.String2Int(videoIdListTmp)
	}
	return

}

// GetVideoFavoriteList 获取视频点赞列表
func (like *Like) GetVideoFavoriteList(videoId int) (userIdList []int) {
	exist := g.DbVideoLike.SIsMember(g.RedisContext, "videoSet", videoId).Val()
	// 缓存不存在
	if !exist {
		// 1、从like表中查询视频ID的点赞列表
		userIdList = like.GetUserIdListForVideo(videoId)
		// 2、将喜欢视频ID列表缓存到redis中
		g.DbVideoLike.SAdd(g.RedisContext, "videoSet", videoId)
		g.DbVideoLike.Set(g.RedisContext, strconv.Itoa(videoId), userIdList, time.Hour*12) //缓存时间设置？
	} else {
		userIdListTmp := g.DbVideoLike.LRange(g.RedisContext, string(videoId), 0, -1).Val()
		userIdList = utils.String2Int(userIdListTmp)
	}
	return
}

// GetVideosFavoriteCount 获取一组视频点赞数
func (like *Like) GetVideosFavoriteCount(videoId []int) (videoFavoriteCount map[int]int) {
	videoFavoriteCount = map[int]int{}
	for _, id := range videoId {
		videoFavoriteList := like.GetVideoFavoriteList(id)
		videoFavoriteCount[id] = len(videoFavoriteList)
	}
	return
}
