package system

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gvaTemplate/global"
	"gvaTemplate/model/system"
	"gvaTemplate/model/system/request"
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

func (s *SystemUserService) ChangePassword(changePasswordModel *request.ChangePassword) (err error) {
	var user system.SysUser
	res := global.GT_DB.Where("user_id = ?", changePasswordModel.UserId).First(&user)
	if res.Error != nil {
		return errors.New("账号不存在")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(changePasswordModel.Password))
	if err != nil {
		return errors.New("密码错误")
	}
	hashedPassword, ExErr := bcrypt.GenerateFromPassword([]byte(changePasswordModel.NewPassword), bcrypt.DefaultCost)
	if ExErr != nil {
		return ExErr
	}
	user.Password = string(hashedPassword)
	err = global.GT_DB.Save(&user).Error
	return err
}
