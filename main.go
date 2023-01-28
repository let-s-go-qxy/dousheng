package main

import "tiktok/boot"

func main() {
	boot.ViperSetup()
	boot.LoggerSetup()
	boot.MysqlDBSetup()
	boot.RedisSetup()
	boot.ServerSetup()
}
