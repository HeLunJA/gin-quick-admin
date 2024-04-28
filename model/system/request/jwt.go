package request

import (
	"github.com/gofrs/uuid/v5"
	"github.com/golang-jwt/jwt/v5"
)

type BaseClaims struct {
	UUID     uuid.UUID
	UserName string
	NickName string
}

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}
