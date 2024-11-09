package user

import (
	"awesomeProject/src/user/dto"
	"awesomeProject/src/user/persistence"
	"github.com/jaswdr/faker/v2"
	"reflect"
	"testing"
)

func TestCreateUserInDb(t *testing.T) {
	userDao := persistence.NewUserDao()
	fake := faker.New()

	user := dto.NewUser("Pinky Pie 1", fake.Internet().Email(), "123456", "Pony")
	userDao.CreateUser(user)
}

func TestSelectUserById(t *testing.T) {
	userDao := persistence.NewUserDao()
	fake := faker.New()

	user := dto.NewUser("Rainbow Dash", fake.Internet().Email(), "12334", "Pony")
	userDao.CreateUser(user)
	readUser := userDao.GetUserById(user.Id)
	if !reflect.DeepEqual(user, readUser) {
		t.Errorf("User is not equal")
	}
}

func TestSelectUserIsNotExist(t *testing.T) {
	userDao := persistence.NewUserDao()

	readUser := userDao.GetUserById(999999)
	if readUser != nil {
		t.Errorf("User should not be returned")
	}
}
