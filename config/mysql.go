package config

type MysqlConfig struct {
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Config   string `mapstructure:"config" json:"config" yaml:"config"`       // 高级配置
	Dbname   string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`    // 数据库名
	Username string `mapstructure:"username" json:"username" yaml:"username"` // 数据库用户名
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 数据库密码
	Path     string `mapstructure:"path" json:"path" yaml:"path"`
}

func (m *MysqlConfig) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}
