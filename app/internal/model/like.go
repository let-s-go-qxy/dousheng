package model

import (
	"gorm.io/gorm"
	"strconv"
	g "tiktok/app/global"
	"tiktok/utils"
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
	var existsLike Like
	result := g.MysqlDB.Where(map[string]interface{}{"user_id": userId, "video_id": videoId}).First(&existsLike)
	aLike := Like{UserId: userId, VideoId: videoId, Cancel: g.FavoriteAction}
	if result.Error == gorm.ErrRecordNotFound {
		g.MysqlDB.Select("user_id", "video_id", "cancel").Create(&aLike)
	} else {
		like.UpdateLike(userId, videoId, g.FavoriteAction)
	}
}

// UpdateLike DeleteLike 更新一条点赞记录
func (like *Like) UpdateLike(userId int, videoId int, cancel int) {

	g.MysqlDB.Model(like).Where(map[string]interface{}{"user_id": userId, "video_id": videoId}).Updates(map[string]interface{}{
		"cancel": cancel,
	})
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

	// 缓存存在用户点赞
	if g.DbVideoLike.SIsMember(g.RedisContext, "userSet", userId).Val() {
		g.DbVideoLike.LRem(g.RedisContext, string(userId), 1, videoId)
	}
	// 缓存存在视频的点赞记录
	if g.DbVideoLike.SIsMember(g.RedisContext, "videoSet", videoId).Val() {
		g.DbVideoLike.LRem(g.RedisContext, string(videoId), 1, userId)
	}
}

// RefreshLikeCache 定期刷新缓存到数据库,likeAdd 和 likeDelete
func (like *Like) RefreshLikeCache() {

	likeAddMap, _ := g.DbVideoLike.HGetAll(g.RedisContext, "likeAdd").Result()
	likeDeleteMap, _ := g.DbVideoLike.HGetAll(g.RedisContext, "likeDelete").Result()

	for key, value := range likeAddMap {
		g.DbVideoLike.HDel(g.RedisContext, "likeAdd", key)
		userId, _ := strconv.Atoi(key)
		videoId, _ := strconv.Atoi(value)
		like.InsertLike(userId, videoId)
	}

	for key, value := range likeDeleteMap {
		g.DbVideoLike.HDel(g.RedisContext, "likeAdd", key)
		userId, _ := strconv.Atoi(key)
		videoId, _ := strconv.Atoi(value)
		like.UpdateLike(userId, videoId, g.CancelFavoriteAction)
	}
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

// GetUserFavoriteVideoList 获取用户喜爱的视频ID列表
func (like *Like) GetUserFavoriteVideoList(userId int) (videoIdList []int) {
	exist := g.DbVideoLike.SIsMember(g.RedisContext, "userSet", userId).Val()
	// 缓存不存在
	if !exist {
		// 1、从like表中查询出喜欢的视频ID列表
		videoIdList = like.GetFavoriteVideoIdList(userId)

		// 2、将喜欢视频ID列表缓存到redis中
		g.DbVideoLike.SAdd(g.RedisContext, "userSet", userId)
		for _, value := range videoIdList {
			g.DbVideoLike.LPush(g.RedisContext, strconv.Itoa(userId), value)
		}
	} else {
		// 缓存存在
		videoIdListTmp := g.DbVideoLike.LRange(g.RedisContext, strconv.Itoa(userId), 0, -1).Val()
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
		for _, value := range userIdList {
			g.DbVideoLike.LPush(g.RedisContext, strconv.Itoa(videoId), value)
		}
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

// IsLike 是否喜欢该视频
func (like *Like) IsLike() (b bool, err error) {
	b = false
	err = g.MysqlDB.First(like, "user_id = ? AND video_id = ? AND cancel = ?", like.UserId, like.VideoId, 1).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	} else {
		b = true
	}
	return
}
