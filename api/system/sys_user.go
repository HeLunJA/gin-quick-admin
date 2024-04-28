package system

import (
	"github.com/gin-gonic/gin"
	"gvaTemplate/model/system"
	"net/http"
)

type UserApi struct{}

func (u *UserApi) DeleteById(c *gin.Context) {
	var userModel system.SysUser
	userId := c.Param("uuid")
	_, err := userService.DeleteUserById(&userModel, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "删除用户成功"})
}
