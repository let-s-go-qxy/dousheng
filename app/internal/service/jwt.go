package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	g "tiktok/app/global"
	"tiktok/app/internal/model"
	"time"
)

type UserClaims struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Expire int    `json:"expire"`
	jwt.StandardClaims
}

// ParseToken 解析token
func ParseToken(str string) (UserClaims, error) {
	c := new(UserClaims)
	token, err := jwt.ParseWithClaims(str, c, func(token *jwt.Token) (interface{}, error) {
		return []byte(g.Config.Auth.Jwt.SecretKey), nil
	})
	if err != nil {
		return *c, errors.New("token不合法")
	}
	if token.Valid != true {
		return *c, errors.New("token不合法")
	}
	return *c, err
}

// GenerateToken 生成7天的token
func GenerateToken(user model.User) (str string) {
	c := &UserClaims{
		Id:     user.Id,
		Name:   user.Name,
		Expire: int(g.Config.Auth.Jwt.ExpiresTime + time.Now().Unix()), // 7天
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 根据自定义key生成tokenString
	str, _ = claims.SignedString([]byte(g.Config.Auth.Jwt.SecretKey))
	return
}
