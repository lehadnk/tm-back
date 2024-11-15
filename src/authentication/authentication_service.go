package authentication

import (
	"fmt"
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
	userFromDB := authService.userService.GetUserByEmailAndPassword(email, password)
	if userFromDB == nil {
		return nil, "", fmt.Errorf("user not found")
	}

	// @todo this is not working since we don't hash passwords on user creation process
	//hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	//if userFromDB.Password != string(hashedPassword) {
	//	return nil, "", fmt.Errorf("wrong credentials")
	//}

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
