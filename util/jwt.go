package util

import (
	"fishba.top/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey=[]byte("fish")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}


func CreateToken(user model.User) (string, error) {
	claims:=&Claims{
		UserId: user.Id,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12*time.Hour).Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "fish",
			Subject: "user token",
		},
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString,err:=token.SignedString(jwtKey)
	if err !=nil {
		return "", err
	}
	return tokenString ,nil
}

func ParseToken(tokenString string)(*jwt.Token,*Claims,error) {
	claims:=&Claims{}
	token,_ := jwt.ParseWithClaims(tokenString,claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey ,nil
	})
	return token ,claims,nil
}