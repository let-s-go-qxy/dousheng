package video

import (
	"sync"
	g "tiktok/app/global"
	"tiktok/app/internal/model"
	"tiktok/app/internal/service/comment"
	"tiktok/app/internal/service/like"
	"tiktok/app/internal/service/user"
	"tiktok/manifest/ossRelated"
)

func GetVideoFeed(latestTime int64, userID int32) (nextTime int64, videoInfo []model.TheVideoInfo, state int) {
	// state 0:已经没有视频了  1:获取成功  -1:获取失败

	allVideoInfoData, isExist := model.VideoDao.GetVideoFeed(int32(latestTime))

	if !isExist {
		// 已经没有视频了
		return nextTime, videoInfo, 0
	}

	nextTime = int64(allVideoInfoData[len(allVideoInfoData)-1].Time)
	videoInfo = make([]model.TheVideoInfo, len(allVideoInfoData))

	wg := sync.WaitGroup{}
	wg.Add(len(allVideoInfoData))

	for index, videoInfoData := range allVideoInfoData {
		var err error
		go func(index int, videoInfo []model.TheVideoInfo, videoInfoData model.VideoInfo, userID int32) {
			var followerCount, followCount, commentCount, favoriteCount int

			var isFollow, isFavorite bool
			_, followCount, followerCount, _, isFollow, err = user.UserInfo(int(userID), int(videoInfoData.UserID))

			_, commentCount = comment.GetCommentList(int(videoInfoData.VideoID), int(userID))
			favoriteCount = like.VideoFavoriteCount(int(videoInfoData.VideoID))
			isFavorite = like.IsLike(int(userID), int(videoInfoData.VideoID))
			avatarURL := user.GetAvatar(int(videoInfoData.UserID))

			videoInfo[index] = model.TheVideoInfo{
				ID: videoInfoData.VideoID,
				Author: model.AuthorInfo{
					ID:            videoInfoData.UserID,
					Name:          videoInfoData.Username,
					FollowCount:   int(followCount),
					FollowerCount: int(followerCount),
					IsFollow:      isFollow,
					Avatar:        avatarURL,
				},
				PlayUrl:       ossRelated.OSSPreURL + videoInfoData.PlayURL + ".mp4",
				CoverUrl:      ossRelated.OSSPreURL + videoInfoData.CoverURL + ".jpg",
				FavoriteCount: int(favoriteCount),
				CommentCount:  int(commentCount),
				IsFavorite:    isFavorite,
				Title:         videoInfoData.Title,
			}
			wg.Done()
		}(index, videoInfo, videoInfoData, userID)
		if err != nil {
			g.Logger.Info("获取视频信息失败，出错了！")
			return nextTime, videoInfo, -1
		}
	}
	wg.Wait()
	return nextTime, videoInfo, 1
}
