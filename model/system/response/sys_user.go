package response

import "gvaTemplate/model/system"

type UserResponse struct {
	UserId   uint        `json:"userId"`
	Username string      `json:"username"`
	NickName string      `json:"nickName"`
	RoleID   *uint       `json:"roleId"`
	Role     system.Role `json:"roleInfo"`
}

type LoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

type CaptchaResponse struct {
	CaptchaImg string `json:"captchaImg"`
	CaptchaId  string `json:"captchaId"`
}
