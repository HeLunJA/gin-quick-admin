package system

import (
	"github.com/gin-gonic/gin"
	"gvaTemplate/global"
	"gvaTemplate/model/system"
	"gvaTemplate/model/system/request"
	"gvaTemplate/model/system/response"
	"gvaTemplate/utils"
)

type BaseApi struct{}

// Register
// @Summary 注册
// @Description 注册用户
// @Tags Bases
// @Produce   application/json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Param nickName body string false "昵称"
// @Success 200 {object} response.Response{data=response.UserResponse,msg=string}  "返回包括用户信息"
// @Failure 500 {object} response.Response{data=nil,msg=string}
// @Router /base/register [post]
func (u *BaseApi) Register(c *gin.Context) {
	var userModel system.SysUser
	if err := c.ShouldBind(&userModel); err != nil {
		response.Fail(err.Error(), c)
		return
	}
	res, err := userService.Register(&userModel)
	if err != nil {
		response.Fail(err.Error(), c)
		return
	}
	regRes := response.UserResponse{
		UserId:   res.UserId,
		Username: res.Username,
		NickName: res.NickName,
	}
	response.Ok(regRes, "注册成功", c)
}

// Login
// @Summary 登录
// @Description 用户登录
// @Tags Bases
// @Produce   application/json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Success 200 {object} response.Response{data=response.LoginResponse,msg=string}  "返回包括用户信息, token"
// @Failure 500 {object} response.Response{data=nil,msg=string}
// @Router /base/login [post]
func (u *BaseApi) Login(c *gin.Context) {
	var user system.SysUser
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Fail(err.Error(), c)
		return
	}
	res, err := baseService.Login(&user)
	if err != nil {
		response.Fail(err.Error(), c)
		return
	}
	j := &utils.JWT{SigningKey: []byte(global.GT_CONFIG.JWT.SigningKey)}
	claims := request.BaseClaims{
		UserId:   res.UserId,
		Username: res.Password,
		NickName: res.NickName,
	}
	newClaims := j.CreateClaims(claims)
	token, err := j.CreateToken(newClaims)
	if err != nil {
		response.Fail("获取token失败", c)
		return
	}
	c.Request.Header.Set("Authorization", "Bearer "+token)
	c.Header("Authorization", "Bearer "+token)
	loginRes := response.LoginResponse{
		User: response.UserResponse{
			UserId:   res.UserId,
			Username: res.Username,
			NickName: res.NickName,
		},
		Token: token,
	}
	response.Ok(loginRes, "登录成功", c)
}
