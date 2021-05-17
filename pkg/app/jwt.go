/*
 * @Author: Kim
 * @Date: 2021-05-13 18:17:14
 * @LastEditTime: 2021-05-14 11:18:17
 * @LastEditors: Kim
 * @Description:
 * @FilePath: /template_go/pkg/app/jwt.go
 */
package app

import (
	"demo/global"
	"demo/pkg/util"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"secret_key"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}
