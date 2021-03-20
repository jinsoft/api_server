package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"user_server/config"
)

var (
	CONFIG config.Server
	DB     *gorm.DB
	LOG    *zap.Logger
	CFG    *viper.Viper
	REDIS  *redis.Client
)
