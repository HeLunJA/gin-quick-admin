package global

import (
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gvaTemplate/config"
)

var BlackCache = local_cache.NewCache()
var (
	GT_DB     *gorm.DB
	GT_VP     *viper.Viper
	GT_CONFIG config.Config
)
