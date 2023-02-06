package main

import (
	"fmt"
	g "tiktok/app/global"
	"tiktok/boot"
	"time"
)

func main() {
	// 文件读取
	boot.ViperSetup()
	// 打印和日志功能  g.Logger
	boot.LoggerSetup()
	// mysql  g.MysqlDB
	boot.MysqlDBSetup()
	// redis  g.DbVerify  g.DbUserLike  g.DbVideoLike
	boot.RedisSetup()
	// 定时任务开启
	boot.CronTaskSetUp()
	// hertz  8080端口
	go boot.ServerSetup()
	for {
		var cursor uint64
		keys, cursor, err := g.DbUserLike.Scan(g.RedisContext, cursor, "*", 100).Result()
		if err != nil {
			fmt.Println("scan keys failed err:", err)
			return
		}
		for _, key := range keys {
			g.DbUserLike.SMembers(g.RedisContext, key).Result()
			//fmt.Println(key, res)
		}
		time.Sleep(time.Second * 3)
	}
}
