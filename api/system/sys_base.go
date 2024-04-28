package system

import (
	"github.com/gin-gonic/gin"
	"gvaTemplate/global"
	"gvaTemplate/model/system"
	"gvaTemplate/model/system/request"
	"gvaTemplate/utils"
	"net/http"
)

type BaseApi struct{}

func (u *BaseApi) Register(c *gin.Context) {
	var userModel system.SysUser
	if err := c.ShouldBind(&userModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := userService.Register(&userModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "注册成功", "data": gin.H{"userId": res.UserId, "userName": res.Username}})
}

func (u *BaseApi) Login(c *gin.Context) {
	var user system.SysUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := baseService.Login(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "获取token失败"})
		return
	}
	c.Request.Header.Set("Authorization", "Bearer "+token)
	c.Header("Authorization", "Bearer "+token)
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "登录成功", "userId": res.UserId})
}
