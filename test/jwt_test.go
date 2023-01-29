package test

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

type UserClaims struct {
	Name   string `json:"name"`
	Expire int    `json:"expire"`
	jwt.StandardClaims
}

// ParseToken 解析token
func TestParseToken(t *testing.T) {
	str := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWVpIiwiZXhwaXJlIjoxNjc1NTkwODE4fQ.r4vvZJElJCDW5QxHVids16poTqAKj-yiMd4_RLm1JV4"
	c := new(UserClaims)
	token, err := jwt.ParseWithClaims(str, c, func(token *jwt.Token) (interface{}, error) {
		return []byte("123312"), nil
	})
	if token.Valid != true || err != nil {
		t.Fail()
	}
}

func TestGT(t *testing.T) {
	c := &UserClaims{
		Name:   "aei",
		Expire: int(604800 + time.Now().Unix()), // 7天
	}
	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 根据自定义key生成tokenString
	str, err := claim.SignedString([]byte("123312"))
	fmt.Println(str, err)
	if str == "" || err != nil {
		t.Fail()
		return
	}
}
