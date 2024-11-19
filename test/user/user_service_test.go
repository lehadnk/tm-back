package user

import (
	"testing"
	"tm/src/user"
	"tm/src/user/persistence"
)

func TestCreateUser(t *testing.T) {
	userService := user.NewUserService(persistence.NewUserDao())
	name := "test"
	email := "admin@test.com"
	password := "qwe123"
	role := "admin"

	newAdminUser := userService.CreateUser(name, email, password, role)
	if newAdminUser.Name != name {
		t.Fatal("Expected", name, "got", newAdminUser.Name)
	}
}
