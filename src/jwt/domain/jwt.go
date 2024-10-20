package domain

import "github.com/golang-jwt/jwt/v5"

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
	Role     string
}

func GenerateToken(UserId int) (string, error) {

	key := []byte("MSWqLa+iY6OUpof6qBHmbeAEdSJPBKrKScpaA222T5M=")
	token := jwt.New(jwt.SigningMethodHS256)
	return token.SignedString(key)
}
