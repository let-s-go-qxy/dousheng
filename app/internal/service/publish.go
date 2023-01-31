package service

import (
	"tiktok/app/internal/model"
)

// 获取登录用户的视频发布列表
func GetPublicList(userId int) (videoList []model.Video, err error) {

	videoList = model.GetPublicList(userId)

	return
}
