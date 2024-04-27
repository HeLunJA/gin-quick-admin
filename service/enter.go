package service

import "gvaTemplate/service/system"

type ServiceGroup struct {
	SystemServiceGroup system.SystemServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
