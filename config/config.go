package config

type Config struct {
	Mysql     MysqlConfig     `mapstructure:"mysql"`
	Server    ServerConfig    `mapstructure:"server"`
	JWT       JWT             `mapstructure:"jwt"`
	System    SystemConfig    `mapstructure:"system"`
	Captcha   CaptchaConfig   `mapstructure:"captcha"`
	AliyunOSS AliyunOSSConfig `mapstructure:"aliyun-oss"`
}
