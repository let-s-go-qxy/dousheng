package main

import (
	"tiktok/boot"
)

func main() {
	// 文件读取
	boot.ViperSetup()
	// 打印和日志功能  g.Logger
	boot.LoggerSetup()
	// mysql  g.MysqlDB
	boot.MysqlDBSetup()
	// redis  g.DbVerify
	boot.RedisSetup()
	// hertz  8080端口
	boot.ServerSetup()
}
