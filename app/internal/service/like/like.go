package like

import (
	"errors"
	"github.com/jinzhu/copier"
	"strconv"
	g "tiktok/app/global"
	repository "tiktok/app/internal/model"
	"tiktok/app/internal/service/user"
	"tiktok/manifest/ossRelated"
	"time"
)

var (
	like   repository.Like
	follow repository.Follow
)

// FavoriteAction 点赞和取消点赞操作
func FavoriteAction(userId int, videoId int, action int) error {
	strUserId := strconv.Itoa(userId)
	strVideoId := strconv.Itoa(videoId)

	// 点赞操作
	if action == g.FavoriteAction {
		// 查询Redis中是否已缓存过该用户的点赞列表
		// 1、已缓存
		if n, err := g.DbUserLike.Exists(g.RedisContext, strUserId).Result(); n > 0 {
			if err != nil {
				g.Logger.Error("方法FavoriteAction执行失败 %v", err)
				return err
			}
			if _, err1 := g.DbUserLike.SAdd(g.RedisContext, strUserId, videoId).Result(); err != nil {
				g.Logger.Error("方法FavoriteAction执行失败 %v", err)
				return err1
			} else {
				// 将点赞/取消点赞缓存在redis中，以"strUserId:videoId的形式存储"，按照时间顺序，定期更新回数据库
				// 后期替换为消息队列
				g.DbUserLike.LPush(g.RedisContext, "likeAdd", strUserId+":"+strVideoId)
			}
		} else {
			//2 未缓存
			// 从数据库拉取用户的点赞列表,并缓存到redis中中
			videoIdList := like.GetFavoriteVideoIdList(userId)
			for _, value := range videoIdList {
				if _, err := g.DbUserLike.SAdd(g.RedisContext, strUserId, value).Result(); err != nil {
					g.Logger.Error("方法：favoriteAction执行失败 %v", err)
					// 防止脏读
					g.DbUserLike.Del(g.RedisContext, strUserId)
					return err
				}
			}

			if _, err := g.DbUserLike.Expire(g.RedisContext, strUserId, time.Minute*5).Result(); err != nil {
				g.Logger.Error("方法favoriteAction：设置过期时间失败%v", err)
				g.DbUserLike.Del(g.RedisContext, strUserId)
				return err
			}
			//
			if _, err := g.DbUserLike.SAdd(g.RedisContext, strUserId, videoId).Result(); err != nil {
				g.Logger.Error("方法：favoriteAction执行失败 %v", err)
				g.DbUserLike.Del(g.RedisContext, strUserId)
				return err
			} else {
				// 替换消息队列
				g.DbUserLike.LPush(g.RedisContext, "likeAdd", strUserId+":"+strVideoId)
			}
		}

		// 查询当前video的点赞列表是否已缓存
		// 1、已缓存
		if n, err := g.DbVideoLike.Exists(g.RedisContext, strVideoId).Result(); n > 0 {
			if err != nil {
				g.Logger.Error("方法：favoriteAction: 缓存查询video点赞列表执行失败 %v", err)
				return err
			}
			if _, err := g.DbVideoLike.SAdd(g.RedisContext, strVideoId, userId).Result(); err != nil {
				g.Logger.Error("方法favoriteAction: video点赞列表插入user执行失败 %v", err)
				return err
			}
		} else {
			//2、未缓存
			userIdList := like.GetUserIdListForVideo(videoId)
			for _, value := range userIdList {
				if _, err := g.DbVideoLike.SAdd(g.RedisContext, strVideoId, value).Result(); err != nil {
					g.Logger.Error("方法favoriteAction:video点赞列表插入执行失败 %v", err)
					// 防止脏读
					g.DbVideoLike.Del(g.RedisContext, strVideoId)
					return err
				}
			}

			if _, err := g.DbVideoLike.Expire(g.RedisContext, strVideoId, time.Minute*5).Result(); err != nil {
				g.Logger.Error("方法favoriteAction：设置过期时间失败%v", err)
				g.DbVideoLike.Del(g.RedisContext, strVideoId)
				return err
			}
			if _, err := g.DbVideoLike.SAdd(g.RedisContext, strVideoId, userId).Result(); err != nil {
				g.Logger.Error("方法favoriteAction:video点赞插入执行失败 %v", err)
				// 防止脏读
				g.DbVideoLike.Del(g.RedisContext, strVideoId)
				return err
			}
		}

		//like.InsertLike(userId, videoId)
	} else if action == g.RequestCancelFavoriteAction { //取消点赞操作
		//缓存存在用户喜爱列表
		if n, err := g.DbUserLike.Exists(g.RedisContext, strUserId).Result(); n > 0 {
			if err != nil {
				g.Logger.Error("方法favoriteAction:缓存查询用户ID执行失败 %v", err)
				return err
			}
			if _, err1 := g.DbUserLike.SRem(g.RedisContext, strUserId, videoId).Result(); err1 != nil {
				g.Logger.Error("方法favoriteAction:缓存取消点赞执行失败 %v", err)
				return err1
			} else {
				// 后期替换消息队列
				g.DbUserLike.LPush(g.RedisContext, "likeDel", strUserId+":"+strVideoId)
			}
		} else { //缓存不存在用户喜爱列表
			// 从数据库拉取最新的点赞列表,并缓存到数据库中
			videoIdList := like.GetFavoriteVideoIdList(userId)
			for _, value := range videoIdList {
				if _, err := g.DbUserLike.SAdd(g.RedisContext, strUserId, value).Result(); err != nil {
					g.Logger.Error("方法：favoriteAction取消点赞执行失败 %v", err)
					// 防止脏读
					g.DbUserLike.Del(g.RedisContext, strUserId)
					return err
				}
			}
			if _, err := g.DbUserLike.Expire(g.RedisContext, strUserId, time.Minute*5).Result(); err != nil {
				g.Logger.Error("方法favoriteAction：设置过期时间失败%v", err)
				g.DbUserLike.Del(g.RedisContext, strUserId)
				return err
			}
			if _, err := g.DbUserLike.LRem(g.RedisContext, strUserId, 1, videoId).Result(); err != nil {
				g.Logger.Error("方法：favoriteAction缓存取消点赞执行失败 %v", err)
				return err
			} else {
				// 替换消息队列
				g.DbUserLike.LPush(g.RedisContext, "likeDel", strUserId+":"+strVideoId)
			}
		}

		// 查询当前video的点赞列表是否已缓存
		// 1、已缓存
		if n, err := g.DbVideoLike.Exists(g.RedisContext, strVideoId).Result(); n > 0 {
			if err != nil {
				g.Logger.Error("方法：favoriteAction: 缓存查询video点赞列表执行失败 %v", err)
				return err
			}
			if _, err := g.DbVideoLike.SRem(g.RedisContext, strVideoId, userId).Result(); err != nil {
				g.Logger.Error("方法favoriteAction: video取消点赞执行失败 %v", err)
				return err
			}
		} else {
			//2、未缓存
			// 从MySQl拉取视频的点赞列表，加载到redis中
			userIdList := like.GetUserIdListForVideo(videoId)
			for _, value := range userIdList {
				if _, err := g.DbVideoLike.SAdd(g.RedisContext, strVideoId, value).Result(); err != nil {
					g.Logger.Error("方法favoriteAction:video取消点赞执行失败 %v", err)
					// 防止脏读
					g.DbVideoLike.Del(g.RedisContext, strVideoId)
					return err
				}
			}
			if _, err := g.DbVideoLike.Expire(g.RedisContext, strVideoId, time.Minute*5).Result(); err != nil {
				g.Logger.Error("方法favoriteAction：设置过期时间失败%v", err)
				g.DbVideoLike.Del(g.RedisContext, strVideoId)
				return err
			}

			if _, err := g.DbVideoLike.SRem(g.RedisContext, strVideoId, userId).Result(); err != nil {
				g.Logger.Error("方法favoriteAction:video点赞插入执行失败 %v", err)
				// 防止脏读
				g.DbVideoLike.Del(g.RedisContext, strVideoId)
				return err
			}
		}
	} else {
		//点赞参数不对，错误处理
		return errors.New("Favorite action type is wrong ")
	}
	return nil
}

// GetVideoListByIdList 根据视频ID列表查询视频列表,按照点赞时间顺序
func GetVideoListByIdList(videoIdList []int) (videoList []repository.Video) {
	for _, videoId := range videoIdList {
		video := repository.Video{}
		g.MysqlDB.Table("videos").Where("id = ?", videoId).Take(&video)
		video.CoverUrl = ossRelated.OSSPreURL + video.CoverUrl + ".jpg"
		video.PlayUrl = ossRelated.OSSPreURL + video.PlayUrl + ".mp4"
		videoList = append(videoList, video)
	}
	return
}

// GetFavoriteList 根据用户ID查询用户的喜欢视频列表
func GetFavoriteList(userId int) ([]repository.RespVideo, error) {
	// 用户喜欢的视频ID列表
	videoIdList, err := like.GetFavoriteVideoList(userId)

	if err != nil {
		return nil, err
	}
	// 每个视频点赞数
	videoFavoriteCount, err1 := like.GetVideosFavoriteCount(videoIdList)
	if err1 != nil {
		return nil, err
	}
	// 视频列表
	videoList := GetVideoListByIdList(videoIdList)
	// 视频对应的发布者
	videosAuthor := GetVideosAuthor(userId, videoList)

	var respVideoList []repository.RespVideo

	//封装返回视频RespVideo列表
	for _, video := range videoList {
		respVideo := repository.RespVideo{}
		copier.Copy(&respVideo, &video)
		respVideo.FavoriteCount = videoFavoriteCount[int(video.Id)]
		respVideo.Author = videosAuthor[int(video.Id)]
		respVideo.IsFavorite = true
		respVideoList = append(respVideoList, respVideo)
	}
	return respVideoList, nil
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
		author.Avatar = user.GetAvatar(author.Id)
		videosAuthor[int(video.Id)] = author
	}
	return
}

// FavoriteVideoCount 根据用户ID查询用户喜爱的视频数目
func FavoriteVideoCount(userId int) int {
	count, _ := like.GetFavoriteVideoList(userId)
	return len(count)
}

// VideoFavoriteCount 根据视频ID查看视频点赞数
func VideoFavoriteCount(videoId int) int {
	userIdList, _ := like.GetVideoFavoriteList(videoId)
	return len(userIdList)
}

// IsLike 根据userId查询用户是否喜欢Id为videoId的视频
func IsLike(userId, videoId int) (b bool) {
	//like.VideoId = videoId
	b, _ = like.IsLike(userId, videoId)
	//fmt.Println(userId, videoId, b)
	return
}
