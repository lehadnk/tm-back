package persistence

import (
	"awesomeProject/src/user/domain"
	"awesomeProject/src/user/persistence"
	"github.com/jaswdr/faker/v2"
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
	if user.Id != readUser.Id {
		t.Errorf("Ids are not matching")
	}
	if user.Name != readUser.Name {
		t.Errorf("Names are not matching")
	}
	if user.Email != readUser.Email {
		t.Errorf("Emails are not matching")
	}
	if user.Password != readUser.Password {
		t.Errorf("Passwords are not matching")
	}
	if user.Role != readUser.Role {
		t.Errorf("Roles are not matching")
	}
}

func TestSelectUserIsNotExist(t *testing.T) {
	userDao := persistence.NewUserDao()

	readUser := userDao.GetUserById(999999)
	if readUser != nil {
		t.Errorf("User should not be returned")
	}
}
