package initialize

import (
	"github.com/gin-gonic/gin"
	"gvaTemplate/global"
)

func RunServer(router *gin.Engine) {
	router.Run(":" + global.GT_CONFIG.Server.Port)
}
