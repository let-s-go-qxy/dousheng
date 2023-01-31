package test


import (
	"fmt"
	"testing"
	g "tiktok/app/global"
)

type Video struct {
	Id          int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Author      int32  `gorm:"column:author_id;not null" json:"author"`
	PlayUrl     string `gorm:"column:play_url;not null" json:"play_url"`
	CoverUrl    string `gorm:"column:cover_url;not null" json:"cover_url"`
	PublishTime int32  `gorm:"column:publish_time;not null" json:"time"`
	Title       string `gorm:"column:title;not null" json:"title"`
}

func TestPublicList(t *testing.T){
	var videoList []int
	g.MysqlDB.Table("videos").Select("id").
		Where("author_id= ? ", 1).
		Scan(&videoList)
	
	for _,video := range videoList{
		fmt.Println(video)
	}
		
}
