package dto

type AuthenticationResult struct {
	isSuccess    bool
	errorMessage string
	jwtToken     string
}
