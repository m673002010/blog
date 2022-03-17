package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// token里面补充用户信息
type JWTClaims struct {
	jwt.StandardClaims
	Id       int
	Username string
	Account  string
}

// 盐值
var (
	jwtKey = []byte("blog")
)

// 生成token
func (JWTClaims) GetToken(id int, username string, account string) (string, error) {
	// token的有效期是七天
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	jwtClaims := &JWTClaims{
		Id:       id,
		Username: username,
		Account:  account,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // token有效期
			IssuedAt:  time.Now().Unix(),     // token发放的时间
		},
	}

	// 根据自定义jwt秘钥生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	tokenString, err := token.SignedString(jwtKey)

	// 返回生成错误
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 解析token
func (JWTClaims) ParseToken(tokenString string) (*jwt.Token, *JWTClaims, error) {
	jwtClaims := &JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenString, jwtClaims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, jwtClaims, err
}
