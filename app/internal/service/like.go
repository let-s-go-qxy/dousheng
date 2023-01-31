package service

import (
	"strconv"
	g "tiktok/app/global"
	repository "tiktok/app/internal/model"
	"time"
)

// GetFavoriteList 查询用户喜欢视频列表,及每个视频的点赞数
func GetFavoriteList(userId int) (videoList []repository.Video, videoFavoriteCount map[int]int) {
	var videoIdList []int
	exist := g.DbVideoLike.SIsMember(g.RedisContext, "userSet", userId).Val()
	//缓存不存在
	if !exist {
		// 1、从like表中查询出喜欢的视频ID列表
		videoIdList = repository.GetFavoriteVideoIdList(userId)
		// 2、将喜欢视频ID列表缓存到redis中
		g.DbVideoLike.SAdd(g.RedisContext, "userSet", userId)
		g.DbVideoLike.Set(g.RedisContext, strconv.Itoa(userId), videoIdList, time.Hour*24) //缓存时间设置？
	}
	// 缓存存在,获取视频缓存Id列表
	videoIdList = repository.GetFavoriteVideoIdList(userId)
	// 获取每个视频点赞数
	videoFavoriteCount = repository.GetVideoFavoriteCount(videoIdList)
	//从数据库获取视频列表
	videoList = GetVideoListByIdList(videoIdList)
	return
}
