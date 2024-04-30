package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gvaTemplate/docs"
	"gvaTemplate/global"
	"gvaTemplate/middleware"
	"gvaTemplate/router"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.New()
	docs.SwaggerInfo.BasePath = global.GT_CONFIG.System.RouterPrefix
	Router.GET(global.GT_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	systemRouter := router.RouterGroupApp.System
	baseRouterGroup := Router.Group(global.GT_CONFIG.System.RouterPrefix)
	{
		systemRouter.InitBaseRouter(baseRouterGroup) // 不需要鉴权的路由
	}
	AuthRouterGroup := Router.Group(global.GT_CONFIG.System.RouterPrefix)
	AuthRouterGroup.Use(middleware.JWTAuth())
	{
		systemRouter.InitUserRouter(AuthRouterGroup)
		systemRouter.InitUploadRouter(AuthRouterGroup)
	}
	// 捕获未匹配的路由
	Router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "未知路由"})
	})
	return Router
}
