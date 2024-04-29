package system

import (
	"github.com/gin-gonic/gin"
	"gvaTemplate/model/system"
	"gvaTemplate/model/system/request"
	"net/http"
)

type UserApi struct{}

func (u *UserApi) DeleteById(c *gin.Context) {
	var userModel system.SysUser
	userId := c.Param("userId")
	_, err := userService.DeleteUserById(&userModel, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "删除用户成功"})
}

func (u *UserApi) ChangePassword(c *gin.Context) {
	claims, flag := c.Get("claims")
	if !flag {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Object not found"})
		return
	}
	exClaims, ok := claims.(request.BaseClaims)
	if !ok {
		c.JSON(500, gin.H{"error": "Invalid object type"})
		return
	}
	var changePasswordModel request.ChangePassword
	if err := c.ShouldBindJSON(&changePasswordModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	changePasswordModel.UserId = exClaims.UserId
	err := userService.ChangePassword(&changePasswordModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "修改密码成功"})
}
