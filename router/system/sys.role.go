package system

import (
	"github.com/gin-gonic/gin"
	v "gvaTemplate/api"
)

type roleRouter struct{}

func (s *roleRouter) InitRoleRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	roleRouter := Router.Group("role")
	RoleApi := v.ApiGroupApp.SystemApiGroup.RoleApi
	{
		roleRouter.POST("addRole", RoleApi.AddRole)
	}
	return roleRouter
}
