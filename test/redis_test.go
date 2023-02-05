package test

import (
	"fmt"
	"testing"
	g "tiktok/app/global"
)

func TestRedis(t *testing.T) {

	var cursor uint64
	keys, cursor, err := g.DbVideoLike.Scan(g.RedisContext, cursor, "*", 100).Result()
	if err != nil {
		fmt.Println("scan keys failed err:", err)
		return
	}
	for _, key := range keys {
		//fmt.Println("key:",key)
		sType, err := g.DbVideoLike.Type(g.RedisContext, key).Result()
		if err != nil {
			fmt.Println("get type failed :", err)
			return
		}
		fmt.Printf("key :%v ,type is %v\n", key, sType)
		if sType == "string" {
			val, err := g.DbVideoLike.Get(g.RedisContext, key).Result()
			if err != nil {
				fmt.Println("get key values failed err:", err)
				return
			}
			fmt.Printf("key :%v ,value :%v\n", key, val)
		} else if sType == "list" {
			val, err := g.DbVideoLike.LPop(g.RedisContext, key).Result()
			if err != nil {
				fmt.Println("get list value failed :", err)
				return
			}
			fmt.Printf("key:%v value:%v\n", key, val)
		}
	}
}
