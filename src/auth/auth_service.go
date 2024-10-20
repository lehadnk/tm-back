package auth

import "awesomeProject/src/auth/dto"

type AuthService struct {}


func (authService *AuthService) Login(email string, password string) dto.AuthenticationResult {
	return new dto.AuthenticationResult{}
}