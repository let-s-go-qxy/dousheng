package api

import (
	"tiktok/app/internal/service/like"
)

var (
	LikeCacheToDBJob = like.LikeCacheToDBJob{Name: "LikeRefresh"}
)
