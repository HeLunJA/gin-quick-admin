package system

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gvaTemplate/global"
	"gvaTemplate/model/system"
)

type SystemBaseService struct{}

func (s *SystemUserService) Register(userModel *system.User) (userInter *system.User, err error) {
	res := global.GT_DB.Where("username = ?", userModel.Username).First(&userInter)
	if res.Error == nil {
		return userModel, errors.New("账号已存在")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userModel.Password), bcrypt.DefaultCost)
	userModel.Password = string(hashedPassword)
	result := global.GT_DB.Create(userModel)
	if result.Error != nil {
		err = result.Error
	}
	return userModel, err
}

func (s *SystemBaseService) Login(userModel *system.User) (userInter *system.User, err error) {
	result := global.GT_DB.Where("username = ?", userModel.Username).First(&userInter)
	err = bcrypt.CompareHashAndPassword([]byte(userInter.Password), []byte(userModel.Password))
	if result.Error != nil || err != nil {
		err = errors.New("用户名或密码错误")
	}
	return userInter, err
}
