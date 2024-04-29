package system

import (
	"github.com/gin-gonic/gin"
	v "gvaTemplate/api"
)

type baseRouter struct{}

func (s *baseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	baseRouter := Router.Group("base")
	BaseApi := v.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("register", BaseApi.Register)
		baseRouter.POST("login", BaseApi.Login)
		baseRouter.GET("captcha", BaseApi.Captcha)
	}
	return baseRouter
}
