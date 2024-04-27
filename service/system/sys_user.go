package system

import (
	"github.com/gofrs/uuid/v5"
	"golang.org/x/crypto/bcrypt"
	"gvaTemplate/global"
	model "gvaTemplate/model/system"
)

type SystemUserService struct{}

func (s *SystemUserService) Login() {

}
func (s *SystemUserService) Create(userModel *model.SysUser) (userInter *model.SysUser, err error) {
	userId := uuid.Must(uuid.NewV4())
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userModel.Password), bcrypt.DefaultCost)
	userModel.UUID = userId
	userModel.Password = string(hashedPassword)
	result := global.GT_DB.Create(userModel)
	if result.Error != nil {
		err = result.Error
	}
	return userModel, err
}
func (s *SystemUserService) DeleteUserById(userModel *model.SysUser, userId string) (userInter *model.SysUser, err error) {
	res := global.GT_DB.Where("uuid = ?", userId).First(userModel)
	if res.Error != nil {
		err = res.Error
	} else {
		res.Delete(userModel)
	}
	return userModel, err
}
