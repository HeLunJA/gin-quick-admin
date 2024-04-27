package config

type Config struct {
	Mysql  MysqlConfig   `mapstructure:"mysql"`
	Server ServiceConfig `mapstructure:"server"`
}
