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

// InsertLike 插入一条点赞记录
func (like *Like) InsertLike(userId int, videoId int) {
	aLike := Like{UserId: userId, VideoId: videoId, Cancel: g.FavoriteAction}
	g.MysqlDB.Select("user_id", "video_id", "cancel").Create(&aLike)
}

// DeleteLike 删除一条点赞记录
func (like *Like) DeleteLike(userId int, videoId int) {
	g.MysqlDB.Where(map[string]interface{}{"user_id": userId, "video_id": videoId}).Update("cancel", g.CancelFavoriteAction)
}

// CacheInsertLike InsertLike 插入一条cache点赞记录
func (like *Like) CacheInsertLike(userId int, videoId int) {
	g.DbVideoLike.HSet(g.RedisContext, "likeAdd", userId, videoId)

	//缓存存在该用户点赞记录
	userCacheExist := g.DbVideoLike.SIsMember(g.RedisContext, "userSet", userId).Val()
	if userCacheExist {
		g.DbVideoLike.LPush(g.RedisContext, string(userId), videoId)
	}

	// 缓存存在视频的点赞列表
	videoCacheExist := g.DbVideoLike.SIsMember(g.RedisContext, "videoSet", userId).Val()
	if videoCacheExist {
		g.DbVideoLike.LPush(g.RedisContext, string(videoId), userId)
	}
}

// CacheDeleteLike DeleteLike 删除一条cache点赞记录
func (like *Like) CacheDeleteLike(userId int, videoId int) {
	g.DbVideoLike.HSet(g.RedisContext, "likeDelete", userId, videoId)
	g.DbVideoLike.LRem(g.RedisContext, string(userId), 1, videoId)
	g.DbVideoLike.LRem(g.RedisContext, string(videoId), 1, userId)
}

// RefreshLikeCache 定期刷新缓存到数据库,likeAdd 和 likeDelete
func (like *Like) RefreshLikeCache() {

}

// GetFavoriteVideoIdList 获取用户点赞的视频ID列表
func (like *Like) GetFavoriteVideoIdList(userId int) (videoList []int) {
	g.MysqlDB.Table("likes").Select("video_id").Where("user_id = ? and cancel = ?", userId, g.FavoriteAction).Find(&videoList)
	return
}

// GetUserIdListForVideo 获取视频的点赞列表
func (like *Like) GetUserIdListForVideo(videoId int) (userIdList []int) {
	g.MysqlDB.Table("likes").Select("user_id").Where("video_id = ? and cancel = ?", videoId, g.FavoriteAction).Find(&userIdList)
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
