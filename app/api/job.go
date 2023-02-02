package api

import "tiktok/app/internal/service"

var (
	LikeCacheToDBJob = service.LikeCacheToDBJob{Name: "LikeRefresh"}
)
