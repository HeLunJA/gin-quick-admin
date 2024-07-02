package system

import (
	"gorm.io/gorm"
	"time"
)

type RoleModel struct {
	ID            uint         `json:"id" gorm:"not null;unique;primary_key;comment:角色ID;size:90"`
	RoleName      string       `json:"roleName" gorm:"not null;comment:角色名称;"`
	Description   *string      `json:"description" gorm:"comment:角色描述;"`
	ParentId      *uint        `json:"parentId" gorm:"comment:父级ID;"`
	ChildrenRoles []*RoleModel `json:"childrenRoles" gorm:"-"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func (RoleModel) TableName() string {
	return "sys_roles"
}
