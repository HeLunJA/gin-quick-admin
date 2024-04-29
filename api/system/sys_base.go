package system

import (
	"github.com/gin-gonic/gin"
	"gvaTemplate/global"
	"gvaTemplate/model/system"
	"gvaTemplate/model/system/request"
	"gvaTemplate/utils"
)

type BaseApi struct{}

func (u *BaseApi) Register(c *gin.Context) {
	var userModel system.SysUser
	if err := c.ShouldBind(&userModel); err != nil {
		utils.Fail(err.Error(), c)
		return
	}
	res, err := userService.Register(&userModel)
	if err != nil {
		utils.Fail(err.Error(), c)
		return
	}
	utils.Ok(gin.H{"userId": res.UserId, "userName": res.Username}, "注册成功", c)
}

func (u *BaseApi) Login(c *gin.Context) {
	var user system.SysUser
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Fail(err.Error(), c)
		return
	}
	res, err := baseService.Login(&user)
	if err != nil {
		utils.Fail(err.Error(), c)
		return
	}
	j := &utils.JWT{SigningKey: []byte(global.GT_CONFIG.JWT.SigningKey)}
	claims := request.BaseClaims{
		UserId:   res.UserId,
		UserName: res.Password,
		NickName: res.NickName,
	}
	newClaims := j.CreateClaims(claims)
	token, err := j.CreateToken(newClaims)
	if err != nil {
		utils.Fail("获取token失败", c)
		return
	}
	c.Request.Header.Set("Authorization", "Bearer "+token)
	c.Header("Authorization", "Bearer "+token)
	utils.Ok(gin.H{"userId": res.UserId, "userName": res.Username}, "登录成功", c)
}
