package authentication

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	jwtdomain "tm/src/authentication/domain"
	"tm/src/user"
	userdomain "tm/src/user/dto"
)

type AuthService struct {
	userService *user.UserService
	jwtManager  *jwtdomain.JwtManager
}

func NewAuthService(userService *user.UserService, jwtManager *jwtdomain.JwtManager) *AuthService {
	var newAuthService = AuthService{
		userService,
		jwtManager,
	}
	return &newAuthService
}

func (authService *AuthService) Login(email string, password string) (*userdomain.User, string, error) {
	userFromDB := authService.userService.GetUserByEmail(email)
	if userFromDB == nil {
		return nil, "", fmt.Errorf("user not found")
	}

	if bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(password)) != nil {
		return nil, "", fmt.Errorf("user not found")
	}

	token, err := authService.jwtManager.GenerateToken(userFromDB)
	if err != nil {
		return nil, "", fmt.Errorf("could not generate token")
	}
	return userFromDB, token, nil
}

func (authService *AuthService) GetCurrentUser(token string) (*userdomain.User, error) {
	currentUser, err := authService.jwtManager.ExchangeToken(token)
	if err != nil {
		return nil, fmt.Errorf("could not verify token")
	}
	return currentUser, nil
}
