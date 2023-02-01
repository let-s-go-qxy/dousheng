package model

import (
	"sync"
	g "tiktok/app/global"
)

const TableNameVideo = "video"

// Video mapped from table <video>
type Video struct {
	Id          int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Author      int32  `gorm:"column:author_id;not null" json:"author"` //用户Id author_id
	PlayUrl     string `gorm:"column:play_url;not null" json:"play_url"`
	CoverUrl    string `gorm:"column:cover_url;not null" json:"cover_url"`
	PublishTime int32  `gorm:"column:publish_time;not null" json:"time"`
	Title       string `gorm:"column:title;not null" json:"title"`
}

// RespVideo 喜爱的视频返回模型
type RespVideo struct {
	Id            int    `json:"id,omitempty"`
	Author        Author `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int    `json:"favorite_count,omitempty"`
	CommentCount  int    `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title,omitempty"`
}

// TableName Video's table name
func (*Video) TableName() string {
	return TableNameVideo
}

type VideoInfo struct {
	VideoID       int32
	UserID        int32
	Username      string
	PlayURL       string
	CoverURL      string
	FavoriteCount int
	IsFavorite    bool
	Time          int32
	Title         string
}

type VideoDaoStruct struct {
}

var (
	VideoDao  *VideoDaoStruct
	videoOnce sync.Once
)

func init() {
	videoOnce.Do(func() {
		VideoDao = &VideoDaoStruct{}
	})
}

func (*VideoDaoStruct) PublishVideo(userID int, title string, videoNumID string) bool {
	video := Video{
		Author:   int32(userID),
		PlayUrl:  videoNumID,
		CoverUrl: videoNumID,
		Title:    title,
	}
	g.MysqlDB.Table("videos").Debug().Create(&video)
	return true

}

func GetPublicList(userId int) (videoList []Video) {

	g.MysqlDB.Table("videos").
		Where("author_id= ? ", userId).
		Scan(&videoList)
	println(videoList)
	return
}
