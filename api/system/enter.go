package system

import (
	"gvaTemplate/service"
)

type ApiGroup struct {
	UserApi
	BaseApi
	RoleApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.SystemUserService
	baseService = service.ServiceGroupApp.SystemServiceGroup.SystemBaseService
	roleService = service.ServiceGroupApp.SystemServiceGroup.SystemRoleService
)
