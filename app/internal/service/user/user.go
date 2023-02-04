package user

import (
	"errors"
	"gorm.io/gorm"
	"strconv"
	g "tiktok/app/global"
	"tiktok/app/internal/model"
	utils2 "tiktok/utils"
	"time"
)

// JudgeNameAndPassword 判断用户名和密码是否符合要求
func JudgeNameAndPassword(name, password string) bool {
	if len(name) > 32 || len(password) > 32 || name == "" || len(password) < 6 {
		return false
	}
	return true
}

func UserLogin(name, password string) (userId int, token string, err error) {
	if JudgeNameAndPassword(name, password) == false {
		err = errors.New("账号或密码不符合要求")
		return
	}
	user := &model.User{
		Name: name,
	}
	err = model.GetUser(user)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.New("账号或密码错误")
			return
		}
		err = errors.New("查询失败: " + err.Error())
		return
	}
	if utils2.GetMd5Str(password+user.Salt) != user.Password {
		err = errors.New("账号或密码错误")
		return
	}
	userId = user.Id
	token = GenerateToken(*user)
	return
}

func UserRegister(name, password string) (userId int, token string, err error) {
	if JudgeNameAndPassword(name, password) == false {
		err = errors.New("账号或密码不符合要求")
		return
	}
	salt := strconv.Itoa(int(time.Now().UnixNano()))
	password = utils2.GetMd5Str(password + salt)
	// 填装数据
	user := &model.User{
		Name:     name,
		Password: password,
		Salt:     salt,
	}
	_, err = model.CreateUser(user)
	// 创建失败
	if err != nil {
		if err == g.ErrDbCreateUniqueKeyRepeatedly {
			err = errors.New("User already exist")
			return
		}
		err = errors.New("创建用户失败: " + err.Error())
		return
	}
	userId = user.Id
	token = GenerateToken(*user)
	return
}

// UserInfo 获取用户信息详情 userId为当前看的主页的用户id，myId是根据token判断而来的当前登录的user_id
func UserInfo(myId int, userId int) (Id, FollowCount, FollowerCount int, Name string, IsFollow bool, err error) {
	userDao := &model.User{
		Id: userId,
	}
	err = model.GetUser(userDao)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.New("User not exit")
		}
	}
	Id = userDao.Id
	Name = userDao.Name
	FollowCount = int(model.GetFollowCount(userDao.Id))
	FollowerCount = int(model.GetFollowerCount(userDao.Id))
	IsFollow = model.IsFollow(myId, userId)
	return Id, FollowCount, FollowerCount, Name, IsFollow, err
}
