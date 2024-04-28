package middleware

import (
	"github.com/gin-gonic/gin"
	"gvaTemplate/utils"
	"net/http"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		println(666666666666666)
		// 从请求头中获取 Authorization 字段
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
			c.Abort()
			return
		}
		// 检查 Authorization 字段是否以 "Bearer " 开头
		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			c.Abort()
			return
		}
		// 提取 JWT 令牌
		tokenString := authHeader[7:]
		j := utils.NewJWT()
		claims, err := j.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to get claims"})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
