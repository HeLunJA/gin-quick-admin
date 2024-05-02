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
	if userModel.RoleID != nil {
		var role system.Role
		roleRes := global.GT_DB.Where("id = ?", userModel.RoleID).First(&system.Role{})
		if roleRes.Error != nil {
			return nil, roleRes.Error
		}
		userModel.Roles = append(userModel.Roles, role)
	}
	result := global.GT_DB.Create(userModel)
	if result.Error != nil {
		err = result.Error
	}
	return userModel, err
}

func (s *SystemBaseService) Login(userModel *system.User) (userInter *system.User, err error) {
	result := global.GT_DB.Where("username = ?", userModel.Username).Preload("Roles").Preload("Role").First(&userInter)
	err = bcrypt.CompareHashAndPassword([]byte(userInter.Password), []byte(userModel.Password))
	if result.Error != nil || err != nil {
		return nil, err
	}
	return userInter, err
}
