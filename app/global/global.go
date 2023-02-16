package global

import (
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"tiktok/app/internal/model/config"
)

var (
	Config       *config.Config
	Logger       *zap.SugaredLogger
	MysqlDB      *gorm.DB
	RedisContext = context.Background()
	DbVerify     *redis.Client
	DbUserLike   *redis.Client
	DbVideoLike  *redis.Client
	OssBucket    *oss.Bucket
)
