package system

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"gvaTemplate/global"
	"gvaTemplate/model/system"
	"gvaTemplate/model/system/request"
	"gvaTemplate/model/system/response"
	"gvaTemplate/utils"
)

var store = base64Captcha.DefaultMemStore

type BaseApi struct{}

// Register
// @Summary 注册
// @Description 注册用户
// @Tags Bases
// @Produce   application/json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Param nickName body string false "昵称"
// @Param captcha body string false "验证码"
// @Param captchaId body string false "验证码ID"
// @Success 200 {object} response.Response{data=response.UserResponse,msg=string}  "返回包括用户信息"
// @Failure 500 {object} response.Response{data=nil,msg=string}
// @Router /base/register [post]
func (u *BaseApi) Register(c *gin.Context) {
	var RegisterModel request.CaptchaRegister

	if err := c.ShouldBind(&RegisterModel); err != nil {
		response.Fail(err.Error(), c)
		return
	}
	captcha := RegisterModel.Captcha
	captchaId := RegisterModel.CaptchaId

	if captchaId != "" && captcha != "" && store.Verify(captchaId, captcha, true) {
		user := system.UserModel{
			Username: RegisterModel.Username,
			Password: RegisterModel.Password,
			NickName: RegisterModel.NickName,
			RoleID:   RegisterModel.RoleID,
		}
		res, err := userService.Register(&user)
		if err != nil {
			response.Fail(err.Error(), c)
			return
		}
		regRes := response.UserResponse{
			UserId:   res.ID,
			Username: res.Username,
			NickName: res.NickName,
		}
		response.Ok(regRes, "注册成功", c)
	} else {
		response.Fail("验证码不正确", c)
	}
}

// Login
// @Summary 登录
// @Description 用户登录
// @Tags Bases
// @Produce   application/json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Param captcha body string false "验证码"
// @Param captchaId body string false "验证码ID"
// @Success 200 {object} response.Response{data=response.LoginResponse,msg=string}  "返回包括用户信息, token"
// @Failure 500 {object} response.Response{data=nil,msg=string}
// @Router /base/login [post]
func (u *BaseApi) Login(c *gin.Context) {
	var RegisterModel request.CaptchaRegister

	if err := c.ShouldBind(&RegisterModel); err != nil {
		response.Fail(err.Error(), c)
		return
	}
	captcha := RegisterModel.Captcha
	captchaId := RegisterModel.CaptchaId

	if captchaId != "" && captcha != "" && store.Verify(captchaId, captcha, true) {
		user := system.UserModel{
			Username: RegisterModel.Username,
			Password: RegisterModel.Password,
		}
		res, err := baseService.Login(&user)
		if err != nil {
			response.Fail(err.Error(), c)
			return
		}
		j := &utils.JWT{SigningKey: []byte(global.GT_CONFIG.JWT.SigningKey)}
		claims := request.BaseClaims{
			UserId:   res.ID,
			Username: res.Username,
			NickName: res.NickName,
		}
		newClaims := j.CreateClaims(claims)
		token, tokenErr := j.CreateToken(newClaims)
		if tokenErr != nil {
			response.Fail("获取token失败", c)
			return
		}
		c.Request.Header.Set("Authorization", "Bearer "+token)
		c.Header("Authorization", "Bearer "+token)
		loginRes := response.LoginResponse{
			User: response.UserResponse{
				UserId:   res.ID,
				Username: res.Username,
				NickName: res.NickName,
				Role:     res.Role,
				RoleID:   res.RoleID,
				Roles:    res.Roles,
			},
			Token: token,
		}
		response.Ok(loginRes, "登录成功", c)
	} else {
		response.Fail("验证码不正确", c)
	}
}

// Captcha
// @Summary 获取验证码
// @Description 获取验证码
// @Tags Bases
// @Produce   application/json
// @Success 200 {object} response.Response{data=response.CaptchaResponse,msg=string}
// @Failure 500 {object} response.Response{data=nil,msg=string}
// @Router /base/captcha [get]
func (u *BaseApi) Captcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(global.GT_CONFIG.Captcha.ImgHeight, global.GT_CONFIG.Captcha.ImgWidth, global.GT_CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		response.Fail("验证码获取失败", c)
		return
	}
	captchaResponse := response.CaptchaResponse{
		CaptchaId:  id,
		CaptchaImg: b64s,
	}
	response.Ok(captchaResponse, "获取验证码成功", c)
}
