package persistence

import (
	"awesomeProject/src/user/domain"
	"awesomeProject/src/user/persistence"
	"github.com/jaswdr/faker/v2"
	"reflect"
	"testing"
)

func TestCreateUserInDb(t *testing.T) {
	userDao := persistence.NewUserDao()
	fake := faker.New()

	user := domain.NewUser("Pinky Pie 1", fake.Internet().Email(), "123456", "Pony")
	userDao.CreateUser(user)
}

func TestSelectUserById(t *testing.T) {
	userDao := persistence.NewUserDao()
	fake := faker.New()

	user := domain.NewUser("Rainbow Dash", fake.Internet().Email(), "12334", "Pony")
	userDao.CreateUser(user)
	readUser := userDao.GetUserById(user.Id)
	if reflect.DeepEqual(user, readUser) {
		t.Errorf("User isn't created")
	}
}

func TestSelectUserIsNotExist(t *testing.T) {
	userDao := persistence.NewUserDao()

	readUser := userDao.GetUserById(999999)
	if readUser != nil {
		t.Errorf("User should not be returned")
	}
}
