package main

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key") // @TODO !!! Generate Key !!!

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
