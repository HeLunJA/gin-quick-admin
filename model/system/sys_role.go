package system

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	ID            uint    `json:"id" gorm:"not null;unique;primary_key;comment:角色ID;size:90"`
	RoleName      string  `json:"roleName" gorm:"not null;comment:角色名称;"`
	Description   *string `json:"description" gorm:"not null;comment:角色描述;"`
	ParentId      *uint   `json:"parentId" gorm:"comment:父级ID;"`
	ChildrenRoles []Role  `json:"childrenRoles" gorm:"many2many:role_relations;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func (Role) TableName() string {
	return "sys_roles"
}
