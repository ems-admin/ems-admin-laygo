package config

import (
	"ems-adming-go/internal/model"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// 定义一个密钥用于签名 JWT
var jwtKey = []byte("85c1516741d8bdf9849397ed22d27281")

// Claims 用户结构体
type Claims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Roles    string `json:"roles"`
	jwt.RegisteredClaims
}

// GenerateJWT 生成 JWT
func GenerateJWT(user model.SysUser) (string, error) {

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		ID:       user.Id,
		Username: user.Username,
		Roles:    user.Roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Println(err.Error())
		return "token生成失败", err
	}

	return tokenString, nil
}

// ValidateJWT 验证 JWT
func ValidateJWT(tokenString string) bool {
	if tokenString == "" {
		return false
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return false
	}
	return true
}
