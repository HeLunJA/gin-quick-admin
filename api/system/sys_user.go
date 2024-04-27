package system

import (
	"github.com/gin-gonic/gin"
	"gvaTemplate/model"
	"gvaTemplate/service"
	"net/http"
)

type UserApi struct{}

var userService = service.ServiceGroupApp.SystemServiceGroup
var userModel = model.ModelGroupApp.SystemModelGroup.SysUser

func (u *UserApi) Login(c *gin.Context) {

}

func (u *UserApi) Create(c *gin.Context) {
	if err := c.ShouldBind(&userModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := userService.Create(&userModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "用户注册成功", "data": gin.H{"userId": res.UUID, "userName": res.Username}})
}

func (u *UserApi) DeleteById(c *gin.Context) {
	userId := c.Param("uuid")
	_, err := userService.DeleteUserById(&userModel, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "删除用户成功"})
}
