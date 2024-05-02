package system

import (
	"github.com/gin-gonic/gin"
	"gvaTemplate/model/system"
	"gvaTemplate/model/system/response"
)

type RoleApi struct{}

func (u *RoleApi) AddRole(c *gin.Context) {
	var role system.Role
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
