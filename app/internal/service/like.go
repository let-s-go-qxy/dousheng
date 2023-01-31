package service

import (
	repository "tiktok/app/internal/model"
)

var (
	like   repository.Like
	follow repository.Follow
)

// GetFavoriteList 查询用户喜欢视频列表,及每个视频的点赞数
func GetFavoriteList(userId int) (videoList []repository.Video, videoFavoriteCount map[int]int, videosAuthor map[int]repository.Author) {
	//用户喜欢的视频ID列表
	videoIdList := like.GetFavoriteVideoIdList(userId)
	// 获取每个视频点赞数
	videoFavoriteCount = like.GetVideosFavoriteCount(videoIdList)
	//从数据库获取视频列表
	videoList = GetVideoListByIdList(videoIdList)
	// 获取每个视频对应的发布者
	videosAuthor = GetVideosAuthor(userId, videoList)
	return
}

func GetVideosAuthor(userId int, videoList []repository.Video) (videosAuthor map[int]repository.Author) {
	videosAuthor = map[int]repository.Author{}
	for _, video := range videoList {
		author := repository.Author{}
		author.Id = int(video.AuthorId)
		author.Name = repository.GetNameById(author.Id)
		author.FollowCount = int(follow.GetFollowCount(int(video.AuthorId)))
		author.FollowerCount = int(follow.GetFollowerCount(int(video.AuthorId)))
		author.IsFollow = follow.IsFollowed(userId, int(video.AuthorId))
		videosAuthor[int(video.Id)] = author
	}
	return
}
