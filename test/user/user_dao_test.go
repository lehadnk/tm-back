package user

import (
	"github.com/jaswdr/faker/v2"
	"reflect"
	"testing"
	"tm/src/user/dto"
	"tm/src/user/persistence"
)

func TestCreateUserInDb(t *testing.T) {
	userDao := persistence.NewUserDao()
	fake := faker.New()

	user := dto.NewUser(fake.Person().Name(), fake.Internet().Email(), fake.Internet().Password(), "Pony")
	userDao.CreateUser(user)
}

func TestSelectUserById(t *testing.T) {
	userDao := persistence.NewUserDao()
	fake := faker.New()
	// The test above runs at the same exact epoch, causing faker to initialize with the same seed, therefore the first email given is the same
	fake.Internet().Email()

	user := dto.NewUser(fake.Person().Name(), fake.Internet().Email(), fake.Internet().Password(), "Pony")
	userDao.CreateUser(user)
	readUser := userDao.GetUserById(user.Id)
	if !reflect.DeepEqual(user, readUser) {
		t.Errorf("User and readUser are not equal")
	}
}

func TestSelectUserDoesNotExist(t *testing.T) {
	userDao := persistence.NewUserDao()

	readUser := userDao.GetUserById(999999)
	if readUser != nil {
		t.Errorf("User should be nil")
	}
}
