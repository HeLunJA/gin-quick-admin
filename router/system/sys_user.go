package system

import (
	"github.com/gin-gonic/gin"
	v "gvaTemplate/api"
)

type userRouter struct{}

func (s *userRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	userRouter := Router.Group("user")
	userApi := v.ApiGroupApp.SystemApiGroup
	{
		userRouter.POST("register", userApi.Create)
		userRouter.POST("login", userApi.Login)
		userRouter.DELETE("/deleteUserById/:uuid", userApi.DeleteById)
	}
	return userRouter
}
