package jwt

import domain "awesomeProject/src/jwt/domain"

func GenerateToken(UserId int) (string, error) {
	return domain.GenerateToken(UserId)
}
