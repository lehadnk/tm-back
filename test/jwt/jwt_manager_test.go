package jwt

import (
	jwtdomain "awesomeProject/src/jwt/domain"
	"awesomeProject/src/user"
	userdomain "awesomeProject/src/user/domain"
	"awesomeProject/src/user/persistence"
	"reflect"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	testUser := userdomain.NewUser("Test username", "test@test.com", "qwe123", "Admin")
	userDao := persistence.NewUserDao()
	userDao.CreateUser(testUser)
	userService := user.NewUserService(userDao)
	jwtManager := jwtdomain.NewJwtManager(userService)

	token, err := jwtManager.GenerateToken(testUser)
	if err != nil {
		t.Fatal("Error while generating token" + err.Error())
	}

	var verifiedUser *userdomain.User
	verifiedUser, err = jwtManager.VerifyTokenAndReturnUser(token)
	if err != nil {
		t.Fatal("Error while verifying token" + err.Error())
	}
	if !reflect.DeepEqual(testUser, verifiedUser) {
		t.Fatal("User is not the same")
	}
}
