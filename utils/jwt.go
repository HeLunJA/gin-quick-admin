package utils

import (
	"errors"
	jwt "github.com/golang-jwt/jwt/v5"
	"gvaTemplate/global"
	"gvaTemplate/model/system/request"
	"time"
)

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.GT_CONFIG.JWT.SigningKey),
	}
}

func (j *JWT) CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	bf, _ := ParseDuration(global.GT_CONFIG.JWT.BufferTime)
	ep, _ := ParseDuration(global.GT_CONFIG.JWT.ExpiresTime)
	claims := request.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second), // 缓冲时间
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"GVA"},                   // 受众
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),    // 过期时间
		},
	}
	return claims
}

func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	// 解析 JWT 令牌
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return nil, err
	}

	// 检查 JWT 令牌是否有效
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// 获取解析出的声明
	claims, ok := token.Claims.(*request.CustomClaims)
	if !ok {
		return nil, errors.New("failed to get claims")
	}
	return claims, nil
}
