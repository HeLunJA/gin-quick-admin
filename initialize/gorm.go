package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gvaTemplate/global"
)

func Gorm() *gorm.DB {
	m := global.GT_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		return nil
	} else {
		println("数据库连接成功")
		return db
	}
}
