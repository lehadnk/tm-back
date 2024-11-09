package authentication

import (
	jwtdomain "awesomeProject/src/authentication/domain"
	userdomain "awesomeProject/src/user/dto"
	"awesomeProject/src/user/persistence"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userDao    *persistence.UserDao
	jwtManager *jwtdomain.JwtManager
}

func NewAuthService(userDao *persistence.UserDao, jwtManager *jwtdomain.JwtManager) *AuthService {
	var newAuthService = AuthService{
		userDao,
		jwtManager,
	}
	return &newAuthService
}

func (authService *AuthService) Login(email string, password string) (*userdomain.User, string, error) {
	userFromDB := authService.userDao.GetUserByEmailAndPassword(email, password)
	if userFromDB == nil {
		return nil, "", fmt.Errorf("user not found")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if userFromDB.Password != string(hashedPassword) {
		return nil, "", fmt.Errorf("wrong credentials")
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
