package router

import "gvaTemplate/router/system"

type RouterGroup struct {
	System system.SystemRouterGroup
}

var RouterGroupApp = new(RouterGroup)
