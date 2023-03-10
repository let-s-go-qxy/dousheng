package model

import (
	"gorm.io/gorm"
	"strconv"
	"strings"
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

// InsertLike 插入一条点赞记录
func (like *Like) InsertLike(userId int, videoId int) {
	var existsLike Like
	result := g.MysqlDB.Where(map[string]interface{}{"user_id": userId, "video_id": videoId}).First(&existsLike)
	aLike := Like{UserId: userId, VideoId: videoId, Cancel: g.FavoriteAction}
	// 点赞记录不存在，则插入
	if result.Error == gorm.ErrRecordNotFound {
		g.MysqlDB.Select("user_id", "video_id", "cancel").Create(&aLike)
	} else {
		//点赞记录存在，则更新
		like.UpdateLike(userId, videoId, g.FavoriteAction)
	}
}

// UpdateLike  更新一条点赞记录
func (like *Like) UpdateLike(userId int, videoId int, actionType int) {
	g.MysqlDB.Model(like).Where(map[string]interface{}{"user_id": userId, "video_id": videoId}).Updates(map[string]interface{}{
		"cancel": actionType,
	})
}

// RefreshLikeCache 定期刷新缓存到数据库,likeAdd 和 likeDel
func (like *Like) RefreshLikeCache() {

	likeAdd, _ := g.DbUserLike.LRange(g.RedisContext, "likeAdd", 0, -1).Result()
	likeDel, _ := g.DbUserLike.LRange(g.RedisContext, "likeDel", 0, -1).Result()

	for _, value := range likeAdd {
		g.DbUserLike.LRem(g.RedisContext, "likeAdd", 1, value)
		msg := strings.Split(value, ":")
		userId, _ := strconv.Atoi(msg[0])
		videoId, _ := strconv.Atoi(msg[1])
		like.InsertLike(userId, videoId)
	}

	for _, value := range likeDel {
		g.DbUserLike.LRem(g.RedisContext, "likeDel", 1, value)
		msg := strings.Split(value, ":")
		userId, _ := strconv.Atoi(msg[0])
		videoId, _ := strconv.Atoi(msg[1])
		like.UpdateLike(userId, videoId, g.CancelFavoriteAction)
	}
}

// GetFavoriteVideoIdList 直接查询MySQL，根据userId， 查询用户点赞的视频,返回videoId列表
func (like *Like) GetFavoriteVideoIdList(userId int) (videoIdList []int) {
	g.MysqlDB.Table("likes").Select("video_id").Where(map[string]interface{}{"user_id": userId, "cancel": g.FavoriteAction}).Find(&videoIdList)
	return
}

// GetUserIdListForVideo 直接查询MySQL，根据videoID，从查询视频点赞列表,返回userId列表
func (like *Like) GetUserIdListForVideo(videoId int) (userIdList []int) {
	g.MysqlDB.Table("likes").Select("user_id").Where(map[string]interface{}{"video_id": videoId, "cancel": g.FavoriteAction}).Find(&userIdList)
	return
}

// GetFavoriteVideoList 根据用户ID，获取用户点赞的视频ID列表
// 先请求缓存，缓存不存在再请求数据库
func (like *Like) GetFavoriteVideoList(userId int) ([]int, error) {
	strUserId := strconv.Itoa(userId)
	if n, err := g.DbUserLike.Exists(g.RedisContext, strUserId).Result(); n > 0 { //缓存存在
		if err != nil {
			g.Logger.Error("方法GetUserFavoriteVideoList: 缓存获取用户喜爱列表失败%v", err)
			return nil, err
		}
		if strVideoIdList, err := g.DbUserLike.SMembers(g.RedisContext, strUserId).Result(); err != nil {
			g.Logger.Error("方法GetUserFavoriteVideoList: 缓存获取用户喜爱列表失败%v", err)
			return nil, err
		} else {
			videoIdList := utils.String2Int(strVideoIdList)
			return videoIdList, nil
		}
	} else { //缓存不存在
		// 从数据库查询，并加载到缓存中
		videoIdList := like.GetFavoriteVideoIdList(userId)
		for _, value := range videoIdList {
			if _, err := g.DbUserLike.SAdd(g.RedisContext, strUserId, value).Result(); err != nil {
				g.Logger.Error("方法GetUserFavoriteVideoList: 用户喜爱列表加载入缓存失败%v\", err")
				g.DbUserLike.Del(g.RedisContext, strUserId)
				return nil, err
			}
		}
		if _, err := g.DbUserLike.Expire(g.RedisContext, strUserId, time.Minute*5).Result(); err != nil {
			g.Logger.Error("方法favoriteAction：设置过期时间失败%v", err)
			g.DbUserLike.Del(g.RedisContext, strUserId)
			return nil, err
		}
		return videoIdList, nil

	}
}

// GetVideoFavoriteList 根据视频Id，获取视频点赞列表
// 先查询缓存，再查询数据库
func (like *Like) GetVideoFavoriteList(videoId int) ([]int, error) {
	strVideoId := strconv.Itoa(videoId)
	if n, err := g.DbVideoLike.Exists(g.RedisContext, strVideoId).Result(); n > 0 { //缓存存在
		if err != nil {
			g.Logger.Error("方法GetUserFavoriteVideoList: 缓存获取视频点赞列表失败%v", err)
			return nil, err
		}
		if strUserIdList, err := g.DbVideoLike.SMembers(g.RedisContext, strVideoId).Result(); err != nil {
			g.Logger.Error("方法GetUserFavoriteVideoList: 缓存获取视频点赞列表失败%v", err)
			return nil, err
		} else {
			userIdList := utils.String2Int(strUserIdList)
			return userIdList, nil
		}
	} else { //缓存不存在
		// 从数据库查询，并加载到缓存中
		userIdList := like.GetUserIdListForVideo(videoId)
		for _, value := range userIdList {
			if _, err := g.DbVideoLike.SAdd(g.RedisContext, strVideoId, value).Result(); err != nil {
				g.Logger.Error("方法GetUserFavoriteVideoList: 缓存写入视频点赞列表失败%v", err)
				g.DbVideoLike.Del(g.RedisContext, strVideoId)
				return nil, err
			}
		}
		if _, err := g.DbVideoLike.Expire(g.RedisContext, strVideoId, time.Minute*5).Result(); err != nil {
			g.Logger.Error("方法favoriteAction：设置过期时间失败%v", err)
			g.DbVideoLike.Del(g.RedisContext, strVideoId)
			return nil, err
		}
		return userIdList, nil
	}
}

// GetVideosFavoriteCount 根据一组视频ID获取一组点赞数
func (like *Like) GetVideosFavoriteCount(videoId []int) (map[int]int, error) {
	videoFavoriteCount := map[int]int{}
	for _, id := range videoId {
		videoFavoriteList, err := like.GetVideoFavoriteList(id)
		if err != nil {
			return nil, err
		}
		videoFavoriteCount[id] = len(videoFavoriteList)
	}
	return videoFavoriteCount, nil
}

func (like *Like) IsLike(userId int, videoId int) (b bool, err error) {

	strUserId := strconv.Itoa(userId)
	if n, err := g.DbUserLike.Exists(g.RedisContext, strUserId).Result(); n > 0 { //缓存存在
		if err != nil {
			g.Logger.Error("方法GetUserFavoriteVideoList: 缓存获取用户喜爱列表失败%v", err)
			return false, err
		}
		if res, err := g.DbUserLike.SIsMember(g.RedisContext, strUserId, videoId).Result(); err != nil {
			g.Logger.Error("方法GetUserFavoriteVideoList: 缓存获取用户喜爱列表失败%v", err)
			return false, err
		} else {
			return res, nil
		}
	} else { //缓存不存在
		// 从数据库查询，并加载到缓存中
		videoIdList := like.GetFavoriteVideoIdList(userId)
		for _, value := range videoIdList {
			if _, err := g.DbUserLike.SAdd(g.RedisContext, strUserId, value).Result(); err != nil {
				g.Logger.Error("方法GetUserFavoriteVideoList: 用户喜爱列表加载入缓存失败%v\", err")
				g.DbUserLike.Del(g.RedisContext, strUserId)
				return false, err
			}
		}
		if _, err := g.DbUserLike.Expire(g.RedisContext, strUserId, time.Minute*5).Result(); err != nil {
			g.Logger.Error("方法favoriteAction：设置过期时间失败%v", err)
			g.DbUserLike.Del(g.RedisContext, strUserId)
			return false, err
		}

		if res, err := g.DbUserLike.SIsMember(g.RedisContext, strUserId, videoId).Result(); err != nil {
			g.Logger.Error("方法GetUserFavoriteVideoList: 缓存获取用户喜爱列表失败%v", err)
			return false, err
		} else {
			return res, nil
		}
	}
}
