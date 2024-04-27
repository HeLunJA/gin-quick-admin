package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"gvaTemplate/global"
)

func Viper() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config") // 文件名
	v.AddConfigPath(".")      // 文件路径，这里设置为当前目录
	v.SetConfigType("yaml")
	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Error reading config file: %s\n", err))
	}
	if err := v.Unmarshal(&global.GT_CONFIG); err != nil {
		panic(fmt.Errorf("Error unmarshaling config: %s\n", err))
	}
	return v
}
