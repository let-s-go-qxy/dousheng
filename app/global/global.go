package global

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"tiktok/app/internal/model/config"
)

var (
	Config       *config.Config
	Logger       *zap.SugaredLogger
	MysqlDB      *gorm.DB
	DbVerify     *redis.Client
	RedisContext = context.Background()
)
