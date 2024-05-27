package initialize

import (
	"gvaTemplate/global"
	"gvaTemplate/model/system"
	"os"
)

func RegisterTables() {
	db := global.GT_DB
	err := db.AutoMigrate(
		system.UserModel{},
		system.RoleModel{},
	)
	if err != nil {
		os.Exit(0)
	}
}
