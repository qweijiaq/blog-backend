package utils

import (
	"backend/global"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

type JwtPayload struct {
	//Username string `json:"username"`
	NickName string `json:"nick_name"`
	Role     int    `json:"role"`
	UserId   uint   `json:"user_id"`
	Avatar   string `json:"avatar"`
}

type CustomClaims struct {
	JwtPayload
	jwt.StandardClaims
}

func GenToken(user JwtPayload) (string, error) {
	var MySecret = []byte(global.Config.Jwt.Secret)
	claim := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expires))), // 默认48小时过期
			Issuer:    global.Config.Jwt.Issuer,                                                     // 签发人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}

// 解析 Token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	var MySecret = []byte(global.Config.Jwt.Secret)
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		global.Log.Error(fmt.Sprintf("token parse err: %s", err.Error()))
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
