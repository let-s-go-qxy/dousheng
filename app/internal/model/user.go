package model

type User struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}
