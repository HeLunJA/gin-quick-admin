package initialize

import (
	"github.com/gin-gonic/gin"
	"gvaTemplate/middleware"
	"gvaTemplate/router"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.New()
	systemRouter := router.RouterGroupApp.System
	baseRouterGroup := Router.Group("")
	{
		systemRouter.InitBaseRouter(baseRouterGroup) // 不需要鉴权的路由
	}
	RouterGroup := Router.Group("")
	RouterGroup.Use(middleware.JWTAuth())
	{
		systemRouter.InitUserRouter(RouterGroup)
	}
	// 捕获未匹配的路由
	Router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "未知路由"})
	})
	return Router
}
