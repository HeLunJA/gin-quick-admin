package api

import "gvaTemplate/api/system"

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
