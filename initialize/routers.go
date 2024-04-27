package initialize

import (
	"github.com/gin-gonic/gin"
	"gvaTemplate/router"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.New()
	systemRouter := router.RouterGroupApp.System
	RouterGroup := Router.Group("")
	{
		systemRouter.InitUserRouter(RouterGroup)
	}
	// 捕获未匹配的路由
	Router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "未知路由"})
	})
	return Router
}
