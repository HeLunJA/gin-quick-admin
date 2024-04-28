package request

import (
	"github.com/golang-jwt/jwt/v5"
)

type BaseClaims struct {
	UserId   uint
	UserName string
	NickName string
}

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}
