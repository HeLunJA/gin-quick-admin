package system

import (
	"github.com/gin-gonic/gin"
	"gvaTemplate/model"
	"gvaTemplate/model/system"
	"gvaTemplate/model/system/request"
	"gvaTemplate/model/system/response"
)

type UserApi struct{}

// DeleteById
// @Summary 根据ID删除用户
// @Description 用户删除
// @Tags Users
// @Produce   application/json
// @Param userId path string true "用户ID"
// @Success 200 {object} response.Response{data=nil,msg=string}
// @Failure 500 {object} response.Response{data=nil,msg=string}
// @Router /deleteUserById/{userId} [delete]
func (u *UserApi) DeleteById(c *gin.Context) {
	var userModel system.User
	userId := c.Param("userId")
	_, err := userService.DeleteUserById(&userModel, userId)
	if err != nil {
		response.Fail(err.Error(), c)
		return
	}
	response.Ok(nil, "删除用户成功", c)
}

// ChangePassword
// @Summary 修改密码
// @Description 修改密码
// @Tags Users
// @Produce   application/json
// @Param password body string true "原密码"
// @Param newPassword body string true "新密码"
// @Success 200 {object} response.Response{data=nil,msg=string}
// @Failure 500 {object} response.Response{data=nil,msg=string}
// @Router /user/changePassword [post]
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
	err := userService.ChangePassword(changePasswordModel)
	if err != nil {
		response.Fail(err.Error(), c)
		return
	}
	response.Ok(nil, "修改密码成功", c)
}

// GetUsers
// @Summary 获取用户列表
// @Description 获取用户列表
// @Tags Users
// @Produce   application/json
// @Param page body int true "页码"
// @Param pageSize body int true "条数"
// @Success 200 {object} response.Response{data=response.PageInfoResponse,msg=string}
// @Failure 500 {object} response.Response{data=nil,msg=string}
// @Router /user/getUsers [post]
func (u *UserApi) GetUsers(c *gin.Context) {
	var pageInfo model.PageInfo
	if err := c.ShouldBindJSON(&pageInfo); err != nil {
		response.Fail(err.Error(), c)
		return
	}
	res, err := userService.GetUsers(pageInfo)
	if err != nil {
		response.Fail(err.Error(), c)
		return
	}
	response.Ok(res, "查询成功", c)
}
