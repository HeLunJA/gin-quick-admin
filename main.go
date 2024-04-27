package main

import (
	"gvaTemplate/global"
	"gvaTemplate/initialize"
)

func main() {
	// 读取yaml配置文件
	global.GT_VP = initialize.Viper()
	//连接数据库
	global.GT_DB = initialize.Gorm()
	//初始化路由组
	Router := initialize.Routers()
	if global.GT_DB != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GT_DB.DB()
		defer db.Close()
	}
	//启动服务
	initialize.RunServer(Router)
}
