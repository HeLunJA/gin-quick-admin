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
// @Router /base/addRole [post]
func (s *SystemRoleService) AddRole(role *system.RoleModel) (*system.RoleModel, error) {
	// Check if ParentId is the same as ID
	if role.ParentId != nil && *role.ParentId == role.ID {
		return nil, errors.New("角色ID不能与父级ID相同")
	}

	// Check if the role ID already exists
	var existingRole system.RoleModel
	if err := global.GT_DB.Where("id = ?", role.ID).First(&existingRole).Error; err == nil {
		return nil, errors.New("角色ID已存在")
	}

	// Check if the parent ID exists
	if role.ParentId != nil {
		if err := global.GT_DB.Where("id = ?", role.ParentId).First(&existingRole).Error; err != nil {
			return nil, errors.New("父级ID不存在")
		}
	}

	// Create the new role
	if err := global.GT_DB.Create(role).Error; err != nil {
		return nil, err
	}

	return role, nil
}

// SetRole
// @Summary 修改角色
// @Description 修改角色
// @Tags Roles
// @Produce   application/json
// @Param id body string true "角色ID"
// @Param roleName body string true "角色名称"
// @Param parentId body int false "父级ID"
// @Param description body string false "描述"
// @Success 200 {object} response.Response{data=system.Role,msg=string}  "返回包括角色信息"
// @Failure 500 {object} response.Response{data=nil,msg=string}
// @Router /role/setRole [put]
func (s *SystemRoleService) SetRole(role *system.RoleModel) (roleInter *system.RoleModel, err error) {
	// 检查 ParentId 是否与 ID 相同
	if role.ParentId != nil && *role.ParentId == role.ID {
		return nil, errors.New("角色ID不能与父级ID相同")
	}

	// 定义角色模型变量
	var existingRole, parentRole system.RoleModel

	// 检查角色 ID 是否存在
	if err := global.GT_DB.Where("id = ?", role.ID).First(&existingRole).Error; err != nil {
		return nil, errors.New("角色ID不存在")
	}

	// 检查父级 ID 是否存在
	if role.ParentId != nil {
		if err := global.GT_DB.Where("id = ?", role.ParentId).First(&parentRole).Error; err != nil {
			return nil, errors.New("父级ID不存在")
		}
	}

	// 更新角色
	if err := global.GT_DB.Model(&existingRole).Updates(role).Error; err != nil {
		return nil, err
	}

	return role, nil
}

// GetRoleList
// @Summary 获取角色列表
// @Description 获取角色列表
// @Tags Roles
// @Produce   application/json
// @Success 200 {object} response.Response{data=system.Role,msg=string}  "返回包括角色信息"
// @Failure 500 {object} response.Response{data=nil,msg=string}
// @Router /role/getRoleList [get]
func (s *SystemRoleService) GetRoleList() ([]system.RoleModel, error) {
	// 查询时排除已删除的记录
	var roleList []system.RoleModel
	if err := global.GT_DB.Find(&roleList, "deleted_at IS NULL").Error; err != nil {
		return nil, err
	}

	// 构建角色树
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
		if roles[i].ParentId == nil {
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
