package global

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gvaTemplate/config"
)

var (
	GT_DB     *gorm.DB
	GT_VP     *viper.Viper
	GT_CONFIG config.Config
)
