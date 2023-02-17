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
	// redis  g.DbVerify  g.DbUserLike  g.DbVideoLike
	boot.RedisSetup()
	//阿里OSS 初始化
	boot.OSSInit()
	// 定时任务开启
	boot.CronTaskSetUp()
	// hertz  8080端口
	boot.ServerSetup()
}
