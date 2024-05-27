package system

import (
	"errors"
	"gorm.io/gorm"
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
func (s *SystemRoleService) AddRole(role *system.RoleModel) (roleInter *system.RoleModel, err error) {
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
	var parentModel system.RoleModel
	if role.ParentId != nil {
		parentRes := tx.Where("id = ?", role.ParentId).Take(&parentModel)
		if parentRes.Error != nil {
			err = parentRes.Error
			tx.Rollback()
			return
		}
		parentModel.ChildrenRoles = append(parentModel.ChildrenRoles, role)
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

func (s *SystemRoleService) SetRole(role *system.RoleModel) (roleInter *system.RoleModel, err error) {
	tx := global.GT_DB.Begin()
	var searchRole system.RoleModel
	if err := tx.Take(&searchRole, role.ID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	var parentModel system.RoleModel
	if err := tx.Preload("ChildrenRoles").Take(&parentModel, searchRole.ParentId).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	result := tx.Updates(&role)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	if err := tx.Model(&parentModel).Association("ChildrenRoles").Delete(&parentModel.ChildrenRoles); err != nil {
		tx.Rollback()
		return nil, err
	}
	parentModel.ChildrenRoles = append(parentModel.ChildrenRoles, role)
	if err := tx.Save(&parentModel).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (s *SystemRoleService) GetRoleList() (roles []system.RoleModel, err error) {
	var roleList []system.RoleModel
	// 查询时排除已删除的记录
	if err := global.GT_DB.Find(&roleList, &system.RoleModel{DeletedAt: gorm.DeletedAt{}}).Error; err != nil {
		return nil, err
	}

	tree, err := buildTree(roleList)
	if err != nil {
		return nil, err
	}

	var roleData []system.RoleModel
	roleData = append(roleData, *tree)
	return roleData, nil
}

func buildTree(roles []system.RoleModel) (*system.RoleModel, error) {
	rolesMap := make(map[uint]*system.RoleModel)
	for i := range roles {
		rolesMap[roles[i].ID] = &roles[i]
	}
	var root *system.RoleModel
	for i := range roles {
		if *roles[i].ParentId == 0 {
			root = &roles[i]
		} else {
			parent := rolesMap[*roles[i].ParentId]
			if parent != nil {
				parent.ChildrenRoles = append(parent.ChildrenRoles, &roles[i])
			}
		}
	}
	return root, nil
}
