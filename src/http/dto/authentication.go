package dto

import "tm/src/user/dto"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	IsSuccess           bool      `json:"isSuccess"`
	AuthenticationToken string    `json:"authenticationToken"`
	User                *dto.User `json:"user"`
}

type AuthenticatedUser struct {
	Email string `json:"email"`
	Type  string `json:"type"`
}
