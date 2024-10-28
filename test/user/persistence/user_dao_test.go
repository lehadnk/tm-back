package persistence

import (
	"awesomeProject/src/user/domain"
	"awesomeProject/src/user/persistence"
	"github.com/jaswdr/faker/v2"
	"testing"
)

func TestCreateUserInDb(t *testing.T) {
	conn := persistence.DbConnection{}
	conn.Connect()

	fake := faker.New()

	user := domain.NewUser("Pinky Pie 1", fake.Internet().Email(), "123456", "Pony")
	conn.CreateUser(user)
}

func TestSelectUserById(t *testing.T) {
	conn := persistence.DbConnection{}
	conn.Connect()

	fake := faker.New()

	user := domain.NewUser("Rainbow Dash", fake.Internet().Email(), "12334", "Pony")
	conn.CreateUser(user)
	readUser := conn.GetUserById(user.Id)
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
	conn := persistence.DbConnection{}
	conn.Connect()

	readUser := conn.GetUserById(999999)
	if readUser != nil {
		t.Errorf("User should not be returned")
	}
}
