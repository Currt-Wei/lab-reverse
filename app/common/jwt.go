package common

import (
	"github.com/dgrijalva/jwt-go"
	"lab-reverse/app/model"
	"time"
)

var jwtKey = []byte("2022lab_portal_system")


type Claims struct {
	UserId uint
	Account		string
	Enable		int
	RoleId		int
	Identity	string
	Email 		string
	jwt.StandardClaims
}

// ReleaseToken 生成token
func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)	// 过期时间
	claims := &Claims{
		UserId: user.Id,
		StandardClaims: jwt.StandardClaims {
			ExpiresAt: expirationTime.Unix(),
			IssuedAt: time.Now().Unix(),	// 签发时间
			Issuer: "scut_cs_lab_portal_system",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, err
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}