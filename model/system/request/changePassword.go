package request

type ChangePassword struct {
	UserId      uint
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}
