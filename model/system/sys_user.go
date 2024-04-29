package system

type SysUser struct {
	UserId   uint   `json:"user_id" gorm:"primarykey;comment:用户ID;not null;"`                              // 用户ID
	Username string `json:"username" gorm:"index;comment:用户登录名;not null;" binding:"required,min=4,max=11"` // 用户登录名
	Password string `json:"password"  gorm:"comment:用户登录密码;not null;" binding:"required"`                  // 用户登录密码
	NickName string `json:"nick_name" gorm:"default:系统用户;comment:用户昵称;not null;"`                          // 用户昵称
}

func (SysUser) TableName() string {
	return "sys_users"
}
