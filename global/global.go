package global

import (
	"backend/config"

	"github.com/cc14514/go-geoip2"

	"github.com/go-redis/redis"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm/logger"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config   *config.Config
	DB       *gorm.DB
	Log      *logrus.Logger
	MysqlLog logger.Interface
	Redis    *redis.Client
	ESClient *elastic.Client
	AddrDB   *geoip2.DBReader
)
