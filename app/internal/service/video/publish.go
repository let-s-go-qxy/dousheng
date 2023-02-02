package video

import (
	"tiktok/app/internal/model"

	"github.com/jinzhu/copier"
)

// 获取登录用户的视频发布列表
func GetPublicList(userId int) (respVideoList []model.RespVideo, err error) {

	//获取视频数组
	var videoList []model.Video
	videoList = model.GetPublicList(userId)

	//利用封装函数
	respVideoList = PlusAuthor(userId, videoList)
	return
}

// 将author封装到video
func PlusAuthor(userId int, videoList []model.Video) (respVideoList []model.RespVideo) {

	for _, video := range videoList {
		respVideo := model.RespVideo{}
		copier.Copy(&respVideo, &video)
		author := model.Author{}
		author.Id = int(video.Author)
		author.Name = model.GetNameById(author.Id)
		author.FollowCount = int(model.GetFollowCount(int(video.Author)))
		author.FollowerCount = int(model.GetFollowerCount(int(video.Author)))
		author.IsFollow = model.IsFollow(userId, int(video.Author))
		copier.Copy(&respVideo.Author, &author)

		respVideoList = append(respVideoList, respVideo)
	}
	return
}
