package system

import (
	"github.com/gin-gonic/gin"
	"gvaTemplate/api/response"
	"gvaTemplate/model/system"
	"gvaTemplate/model/system/request"
)

type UserApi struct{}

func (u *UserApi) DeleteById(c *gin.Context) {
	var userModel system.SysUser
	userId := c.Param("userId")
	_, err := userService.DeleteUserById(&userModel, userId)
	if err != nil {
		response.Fail(err.Error(), c)
		return
	}
	response.Ok(nil, "删除用户成功", c)
}

func (u *UserApi) ChangePassword(c *gin.Context) {
	claims, flag := c.Get("claims")
	if !flag {
		response.NoAuth("Object not found", c)
		return
	}
	exClaims, ok := claims.(request.BaseClaims)
	if !ok {
		response.Fail("Invalid object type", c)
		return
	}
	var changePasswordModel request.ChangePassword
	if err := c.ShouldBindJSON(&changePasswordModel); err != nil {
		response.Fail(err.Error(), c)
		return
	}
	changePasswordModel.UserId = exClaims.UserId
	err := userService.ChangePassword(&changePasswordModel)
	if err != nil {
		response.Fail(err.Error(), c)
		return
	}
	response.Ok(nil, "修改密码成功", c)
}
