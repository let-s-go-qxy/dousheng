package model

import (
	"gorm.io/gorm"
	g "tiktok/app/global"
)

type User struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

// CreateUser 增加用户
func CreateUser(user *User) (db *gorm.DB, err error) {
	var count int64
	db = g.MysqlDB.Model(&User{}).Where("name = ?", user.Name).Count(&count)
	if count > 0 {
		err = g.ErrDbCreateUniqueKeyRepeatedly
		return
	}
	db = g.MysqlDB.Create(user)
	err = db.Error
	return
}

func GetUserByName(name string) (user *User, err error) {
	user = new(User)
	err = g.MysqlDB.First(user, "name = ?", name).Error
	return
}
