package auth

import (
	"github.com/jaswdr/faker/v2"
	"reflect"
	"testing"
	jwtdomain "tm/src/authentication/domain"
	"tm/src/user"
	userdomain "tm/src/user/dto"
	"tm/src/user/persistence"
)

func TestGenerateToken(t *testing.T) {
	userDao := persistence.NewUserDao()
	userService := user.NewUserService(userDao)
	jwtManager := jwtdomain.NewJwtManager(userService)

	fake := faker.New()

	testUser := userdomain.NewUser(fake.Person().Name(), fake.Internet().Email()+"1", fake.Internet().Password(), "admin")
	userDao.CreateUser(testUser)

	token, err := jwtManager.GenerateToken(testUser)
	if err != nil {
		t.Fatal("Error while generating token: " + err.Error())
	}

	var verifiedUser *userdomain.User
	verifiedUser, err = jwtManager.ExchangeToken(token)
	if err != nil {
		t.Fatal("Error while verifying token: " + err.Error())
	}

	if !reflect.DeepEqual(testUser, verifiedUser) {
		t.Fatal("testUser and verifiedUser are not the same")
	}
}
