package core

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
)

func (au authCore) NewJwt(user *model.User) string {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte("SECRET_KEY"))

	return tokenString
}
