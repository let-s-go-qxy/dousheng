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

// GetUser 通过名称和user_id查询记录 limit 1
func GetUser(user *User) (err error) {
	if user.Name != "" {
		err = g.MysqlDB.First(user, "name = ?", user.Name).Error
		return
	}
	err = g.MysqlDB.First(user, "id = ?", user.Id).Error
	return
}

func GetNameById(userId int) string {
	var user User
	g.MysqlDB.Where("id = ?", userId).Take(&user)
	return user.Name
}
