package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JWTsecret = []byte("ABAB")

// Claims 是Token所有者的结构
type Claims struct {
	Id       uint   `json:"id"`        // 用户的id
	UserName string `json:"user_name"` // 用户名
	Password string `json:"password"`  // 用户密码
	jwt.StandardClaims
}

// GenerateToken 签发 Token
func GenerateToken(id uint, username, password string) (string, error) {
	// 获取当前时间
	nowTime := time.Now() // 当前登录时间
	// 登录24小时之后的时间
	expireTime := nowTime.Add(24 * time.Hour)

	claims := Claims{
		Id:       id,
		UserName: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "memo",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JWTsecret)
	return token, err
}

// ParseToken 解析 Token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTsecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
