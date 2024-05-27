package system

import (
	"github.com/gin-gonic/gin"
	"gvaTemplate/model/system"
	"gvaTemplate/model/system/response"
)

type RoleApi struct{}

func (u *RoleApi) AddRole(c *gin.Context) {
	var role system.RoleModel
	if err := c.ShouldBindJSON(&role); err != nil {
		response.Fail(err.Error(), c)
		return
	}
	if role.ID == 0 {
		response.Fail("id不能为空", c)
		return
	}
	res, err := roleService.AddRole(&role)
	if err != nil {
		response.Fail(err.Error(), c)
		return
	}
	response.Ok(res, "创建成功", c)
}

func (u *RoleApi) GetRoleList(c *gin.Context) {
	res, err := roleService.GetRoleList()
	if err != nil {
		response.Fail(err.Error(), c)
	}
	response.Ok(res, "查询成功", c)
}

func (u *RoleApi) SetRole(c *gin.Context) {
	var role system.RoleModel
	if err := c.ShouldBindJSON(&role); err != nil {
		response.Fail(err.Error(), c)
		return
	}
	if *role.ParentId == 0 {
		response.Fail("parentId不能为0", c)
	}
	setRole, err := roleService.SetRole(&role)
	if err != nil {
		return
	}
	response.Ok(setRole, "修改成功", c)
}
