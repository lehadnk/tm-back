package src

import (
	"testing"
)

func TestCreateUserInDb(t *testing.T) {
	conn := DbConnection{}
	conn.Connect()

	user := User{Name: "Pinky Pie", Email: "test@test.com", Password: "123456", Role: "Pony"}
	conn.CreateUser(&user)
}
