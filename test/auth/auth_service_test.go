package auth

import (
	"github.com/jaswdr/faker/v2"
	"testing"
	"tm/src/authentication"
	jwtdomain "tm/src/authentication/domain"
	"tm/src/user"
	userdomain "tm/src/user/dto"
	"tm/src/user/persistence"
)

func TestAddUserAndLogin(t *testing.T) {
	userDao := persistence.NewUserDao()
	userService := user.NewUserService(userDao)
	jwtManager := jwtdomain.NewJwtManager(userService)
	authService := authentication.NewAuthService(userService, jwtManager)
	fake := faker.New()

	testUser := userdomain.NewUser(fake.Person().Name(), fake.Internet().Email(), fake.Internet().Password(), "user")
	userService.CreateUser(testUser.Name, testUser.Email, testUser.Password, testUser.Role)

	newUser, token, err := authService.Login(testUser.Email, testUser.Password)
	if newUser == nil {
		t.Fatal("User not found")
	}
	if token == "" {
		t.Fatal("Token not found")
	}
	if err != nil {
		t.Fatal("Error while logging in: " + err.Error())
	}
}
