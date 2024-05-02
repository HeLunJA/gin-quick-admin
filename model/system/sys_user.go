package system

import (
	"gorm.io/gorm"
	"gvaTemplate/model"
)

type User struct {
	model.Model
	Username  string         `json:"username" gorm:"index;not null;comment:用户登录名;" binding:"required,min=4,max=11"` // 用户登录名
	Password  string         `json:"password"  gorm:"not null;comment:用户登录密码;" binding:"required"`                  // 用户登录密码
	NickName  string         `json:"nickName" gorm:"default:匿名用户;not null;comment:用户昵称;"`
	RoleID    uint           `json:"roleId" gorm:"index;comment:用户角色ID"`
	Role      Role           `json:"Role" gorm:"foreignKey:RoleID;references:RoleID;comment:用户角色;"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (User) TableName() string {
	return "sys_users"
}
