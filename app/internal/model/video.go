package model

import (
	"sync"
)

const TableNameVideo = "video"

// Video mapped from table <video>
type Video struct {
	ID          int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AuthorID    int32  `gorm:"column:author_id;not null" json:"author_id"`
	PlayURL     string `gorm:"column:play_url;not null" json:"play_url"`
	CoverURL    string `gorm:"column:cover_url;not null" json:"cover_url"`
	PublishTime int32  `gorm:"column:publish_time;not null" json:"time"`
	Title       string `gorm:"column:title;not null" json:"title"`
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

	return true
}
