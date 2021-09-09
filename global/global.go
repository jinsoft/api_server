package global

import (
	"api_server/config"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	CONFIG config.Server
	DB     *gorm.DB
	LOG    *zap.Logger
	CFG    *viper.Viper
	REDIS  *redis.Client
)
