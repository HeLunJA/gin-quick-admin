package system

import "github.com/gofrs/uuid/v5"

type SysUser struct {
	UUID     uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID;not null;"`                                    // 用户UUID
	Username string    `json:"username" gorm:"index;comment:用户登录名;not null;" binding:"required,min=6,max=11"` // 用户登录名
	Password string    `json:"password"  gorm:"comment:用户登录密码;not null;" binding:"required"`                  // 用户登录密码
	NickName string    `json:"nick_name" gorm:"default:系统用户;comment:用户昵称;not null;"`                          // 用户昵称
}

func (SysUser) TableName() string {
	return "sys_users"
}
