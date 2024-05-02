package system

import (
	"errors"
	"gvaTemplate/global"
	"gvaTemplate/model/system"
)

type SystemRoleService struct{}

// AddRole
// @Summary 新增角色
// @Description 新增角色
// @Tags Roles
// @Produce   application/json
// @Param id body string true "角色ID"
// @Param roleName body string true "角色名称"
// @Param parentId body int false "父级ID"
// @Param description body string false "描述"
// @Success 200 {object} response.Response{data=system.Role,msg=string}  "返回包括角色信息"
// @Failure 500 {object} response.Response{data=nil,msg=string}
// @Router /base/register [post]
func (s *SystemRoleService) AddRole(role *system.Role) (roleInter *system.Role, err error) {
	tx := global.GT_DB.Begin()
	res := tx.Where("id = ?", role.ID).First(&roleInter)
	if res.Error == nil {
		tx.Rollback()
		return nil, errors.New("角色ID已存在")
	}
	res = tx.Create(role)
	if res.Error != nil {
		err = res.Error
		tx.Rollback()
	}
	var parentModel system.Role
	if role.ParentId != nil {
		parentRes := tx.Where("id = ?", role.ParentId).First(&parentModel)
		if parentRes.Error != nil {
			err = parentRes.Error
			tx.Rollback()
			return
		}
		parentModel.ChildrenRoles = append(parentModel.ChildrenRoles, *role)
		saveRes := tx.Save(&parentModel)
		if saveRes.Error != nil {
			err = saveRes.Error
			tx.Rollback()
			return
		}
		if txErr := tx.Commit().Error; txErr != nil {
			tx.Rollback()
		}
	}
	return role, err
}
