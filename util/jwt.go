package util

/*
	JWT
	"github.com/dgrijalva/jwt-go"
*/

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 创建自己的Claims
type JwtClaims struct {
	*jwt.StandardClaims
	//jwt的数据
	Data string
}


var (
	//盐
	secret []byte = []byte("814337bc8e5c98b183a289dd1234536922")
	issuer        = "mange"
)

// CreateJwtToken 生成一个jwttoken
func CreateJwtToken(data string) (signedToken string, err error) {
	expireSec := 3600000
	expireToken := time.Now().Add(time.Second * time.Duration(expireSec)).Unix()
	claims := JwtClaims{
		&jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: expireToken,
			Issuer:    issuer,
		},
		data,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString(secret)
	return
}

// VerifyToken 得到一个JwtToken,然后验证是否合法,防止伪造
func VerifyJwtToken(jwtToken string) bool {
	_, err := jwt.Parse(jwtToken, func(*jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		fmt.Println("解析jwtToken失败.", err)
		return false
	}
	return true
}

// ParseJwtToken 解析token得到是自己创建的Claims
func ParseJwtToken(jwtToken string) (*JwtClaims, error) {
	var jwtclaim = &JwtClaims{}
	_, err := jwt.ParseWithClaims(jwtToken, jwtclaim, func(*jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		fmt.Println("解析jwtToken失败.", err)
		return nil, errors.New("解析jwtToken失败")
	}
	return jwtclaim, nil
}
