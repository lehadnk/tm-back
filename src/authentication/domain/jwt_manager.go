package domain

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"tm/src/user"
	userdomain "tm/src/user/dto"
)

type JwtManager struct {
	secretKey   []byte
	userService *user.UserService
}

func NewJwtManager(userService *user.UserService) *JwtManager {
	var newJwtManager = JwtManager{
		[]byte("secret-key"),
		userService,
	}
	return &newJwtManager
}

func (jwtManager *JwtManager) GenerateToken(user *userdomain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  user.Id,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(jwtManager.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (jwtManager *JwtManager) ExchangeToken(tokenString string) (*userdomain.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtManager.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !token.Valid || !ok {
		return nil, fmt.Errorf("invalid token")
	}

	userId := int(claims["id"].(float64))

	readUser := jwtManager.userService.GetUserById(userId)
	if readUser == nil {
		return nil, fmt.Errorf("user not found")
	}

	return readUser, nil
}
