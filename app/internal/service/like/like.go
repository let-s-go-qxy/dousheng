package like

import (
	"fmt"
	"github.com/jinzhu/copier"
	g "tiktok/app/global"
	repository "tiktok/app/internal/model"
)

var (
	like   repository.Like
	follow repository.Follow
)

// FavoriteAction 点赞和取消点赞操作
func FavoriteAction(userId int, videoId int, action int) bool {
	if action == g.FavoriteAction {
		// 点赞操作
		like.InsertLike(userId, videoId)
		//like.CacheInsertLike(userId, videoId)
		return true
	} else if action == g.RequestCancelFavoriteAction {
		//取消点赞操作
		like.UpdateLike(userId, videoId, g.CancelFavoriteAction)
		//like.CacheDeleteLike(userId, videoId)
		return true
	} else {
		//点赞参数不对，错误处理
		return false
	}
}

// GetVideoListByIdList 根据视频ID列表查询视频列表,按照点赞时间顺序
func GetVideoListByIdList(videoIdList []int) (videoList []repository.Video) {
	for _, videoId := range videoIdList {
		video := repository.Video{}
		g.MysqlDB.Table("videos").Where("id = ?", videoId).Take(&video)
		videoList = append(videoList, video)
	}
	return
}

// GetFavoriteList 查询用户喜欢视频列表,及每个视频的点赞数
func GetFavoriteList(userId int) (respVideoList []repository.RespVideo) {
	//用户喜欢的视频ID列表
	videoIdList := like.GetUserFavoriteVideoList(userId)
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
		author.Id = int(video.Author)
		author.Name = repository.GetNameById(author.Id)
		author.FollowCount = int(repository.GetFollowCount(int(video.Author)))
		author.FollowerCount = int(repository.GetFollowerCount(int(video.Author)))
		author.IsFollow = repository.IsFollow(userId, int(video.Author))
		videosAuthor[int(video.Id)] = author
	}
	return
}

func IsLike(userId, videoId int) (b bool) {
	like := new(repository.Like)
	like.UserId = userId
	like.VideoId = videoId
	b, _ = like.IsLike()
	fmt.Println(userId, videoId, b)
	return
}
