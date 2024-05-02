package initialize

import (
	"gvaTemplate/global"
	"gvaTemplate/model/system"
	"os"
)

func RegisterTables() {
	db := global.GT_DB
	err := db.AutoMigrate(
		system.User{},
		system.Role{},
		system.RoleRelation{},
	)
	if err != nil {
		os.Exit(0)
	}
}
