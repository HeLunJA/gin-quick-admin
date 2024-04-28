package system

import (
	"gvaTemplate/global"
	"gvaTemplate/model/system"
)

type SystemUserService struct{}

func (s *SystemUserService) DeleteUserById(userModel *system.SysUser, userId string) (userInter *system.SysUser, err error) {
	res := global.GT_DB.Where("user_id = ?", userId).First(userModel)
	if res.Error != nil {
		err = res.Error
	} else {
		res.Delete(userModel)
	}
	return userModel, err
}
