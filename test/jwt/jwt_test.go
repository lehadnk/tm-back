package jwt

import (
	jwtdomain "awesomeProject/src/jwt/domain"
	userdomain "awesomeProject/src/user/domain"
	"awesomeProject/src/user/persistence"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	user := userdomain.NewUser("Test username", "test@test.com", "qwe123", "Admin")
	userDao := persistence.NewUserDao()
	userDao.CreateUser(user)

	token, err := jwtdomain.GenerateToken(user)
	if err != nil {
		t.Fatal("Error while generating token" + err.Error())
	}

	_, err = jwtdomain.VerifyToken(token)
	if err != nil {
		t.Fatal("Error while verifying token" + err.Error())
	}
}
