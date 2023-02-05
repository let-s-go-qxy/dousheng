package boot

import (
	"github.com/henrylee2cn/goutil/calendar/cron"
	"tiktok/app/api"
)

func CronTaskSetUp() {
	c := cron.New()
	c.AddJob("0/30 * * * * ?", api.LikeCacheToDBJob)
	go c.Start()
	defer c.Stop()
}
