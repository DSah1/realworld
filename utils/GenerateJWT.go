package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var SecretString = []byte("!@SECRET@!")

func GenerateJWT(id uint) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, _ := token.SignedString(SecretString)

	return t
}
