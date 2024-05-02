package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvaTemplate/global"
	"log"
	"os"
	"time"
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
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 日志输出到标准输出
		logger.Config{
			SlowThreshold: 100 * time.Millisecond, // 慢 SQL 阈值
			LogLevel:      logger.Info,            // 日志级别
			Colorful:      true,                   // 彩色日志输出
		},
	)
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{Logger: dbLogger}); err != nil {
		return nil
	} else {
		println("数据库连接成功")
		return db
	}
}
