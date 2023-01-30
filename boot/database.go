package boot

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	g "tiktok/app/global"
	"time"
)

func MysqlDBSetup() {
	config := g.Config.DataBase.Mysql
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(config.GetDsn()), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		g.Logger.Fatalf("initialize mysql db failed, err: %v", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxIdleTime(10 * time.Second)
	sqlDB.SetConnMaxLifetime(100 * time.Second)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	err = sqlDB.Ping()
	if err != nil {
		g.Logger.Fatalf("connect to mysql db failed, err: %v", err)
	}

	g.Logger.Infof("initialize mysql db successfully")
	g.MysqlDB = db
}

func RedisSetup() {
	config := g.Config.DataBase.Redis
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	verifyDb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Addr, config.Port),
		Username: "",
		Password: config.Password,
		DB:       config.DbVerify,
		PoolSize: 10000,
	})
	_, err := verifyDb.Ping(ctx).Result()
	if err != nil {
		g.Logger.Fatalf("connect to verify redis instance failed, err: %v", err)
	}
	g.DbVerify = verifyDb

	g.Logger.Info("initialize verify redis client successfully")

	userLikeDb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Addr, config.Port),
		Username: "",
		Password: config.Password,
		DB:       config.DbUserLike,
		PoolSize: 10000,
	})
	_, err = userLikeDb.Ping(ctx).Result()
	if err != nil {
		g.Logger.Fatalf("connect to userLike redis instance failed, err: %v", err)
	}
	g.DbUserLike = userLikeDb

	g.Logger.Info("initialize userLike redis client successfully")

	videoLikeDb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Addr, config.Port),
		Username: "",
		Password: config.Password,
		DB:       config.DbVideoLike,
		PoolSize: 10000,
	})
	_, err = videoLikeDb.Ping(ctx).Result()
	if err != nil {
		g.Logger.Fatalf("connect to videoLike redis instance failed, err: %v", err)
	}
	g.DbVideoLike = videoLikeDb

	g.Logger.Info("initialize videoLike redis client successfully")

}
