package service

import (
	"github.com/jinzhu/copier"
	g "tiktok/app/global"
	repository "tiktok/app/internal/model"
)

var (
	like   repository.Like
	follow repository.Follow
)

func FavoriteAction(userId int, videoId int, action int) bool {
	if action == g.FavoriteAction {
		// 点赞操作
		like.CacheInsertLike(userId, videoId)
		return true
	} else if action == g.CancelFavoriteAction {
		//取消点赞操作
		like.CacheDeleteLike(userId, videoId)
		return true
	} else {
		//点赞参数不对，错误处理
		return false
	}
}

// GetFavoriteList 查询用户喜欢视频列表,及每个视频的点赞数
func GetFavoriteList(userId int) (respVideoList []repository.RespVideo) {
	//用户喜欢的视频ID列表
	videoIdList := like.GetFavoriteVideoIdList(userId)
	// 获取每个视频点赞数
	videoFavoriteCount := like.GetVideosFavoriteCount(videoIdList)
	//从数据库获取视频列表
	videoList := GetVideoListByIdList(videoIdList)
	// 获取每个视频对应的发布者
	videosAuthor := GetVideosAuthor(userId, videoList)

	//封装返回视频RespVideo列表
	for _, video := range videoList {
		respVideo := repository.RespVideo{}
		copier.Copy(&respVideo, &video)
		respVideo.FavoriteCount = videoFavoriteCount[int(video.Id)]
		respVideo.Author = videosAuthor[int(video.Id)]
		respVideo.IsFavorite = true
		respVideoList = append(respVideoList, respVideo)
	}

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
