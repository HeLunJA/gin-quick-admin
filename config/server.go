package config

type ServerConfig struct {
	Port string `mapstructure:"port" json:"port" yaml:"port"`
	Host string `mapstructure:"host" json:"host" yaml:"host"`
}
