package test

import (
	"awesomeProject/src"
	"awesomeProject/src/jwt/domain"
	"awesomeProject/src/user/persistence"
	"testing"
)

func TestCreateUserInDb(t *testing.T) {
	conn := persistence.DbConnection{}
	conn.Connect()

	user := domain.User{Name: "Pinky Pie 1", Email: "test2@test.com", Password: "123456", Role: "Pony"}
	conn.CreateUser(&user)
}

func TestCall(t *testing.T) {
	println(src.Call())
}
