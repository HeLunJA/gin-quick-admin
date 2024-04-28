package system

import (
	"gvaTemplate/service"
)

type ApiGroup struct {
	UserApi
	BaseApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.SystemUserService
	baseService = service.ServiceGroupApp.SystemServiceGroup.SystemBaseService
)
