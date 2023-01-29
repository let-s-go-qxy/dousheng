package model



type Video struct {
	Id          int    `gorm:"primaryKey" json:"id"`
	AuthorId    int    `json:"id"`
	PlayUrl     string `json:"playurl"`
	CoverUrl    string `json:"coverurl"`
	PublishTime int    `json:"publishtime"`
	Title       string `json:"title"`
}
