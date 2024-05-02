package system

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	RoleID        uint    `json:"roleId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"`
	RoleName      string  `json:"roleName" gorm:"not null;comment:角色名称;"`
	Description   *string `json:"description" gorm:"not null;comment:角色描述;"`
	ParentId      *uint   `json:"parentId" gorm:"not null;comment:父级ID;"`
	ChildrenRoles []Role  `json:"childrenRoles" gorm:"many2many:role_relations;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func (Role) TableName() string {
	return "sys_roles"
}

type RoleRelation struct {
	ParentID  uint
	ChildID   uint
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (RoleRelation) TableName() string {
	return "sys_role_relation"
}
