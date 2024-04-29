package response

type UserResponse struct {
	UserId   uint   `json:"userId"`
	Username string `json:"username"`
	NickName string `json:"nickName"`
}

type LoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}
