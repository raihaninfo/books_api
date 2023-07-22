package utils

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24 * 2).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		log.Println(err.Error())
	}
	return tokenString
}
