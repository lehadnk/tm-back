package test

import (
	"awesomeProject/src"
	"testing"
)

func TestCreateUserInDb(t *testing.T) {
	conn := src.DbConnection{}
	conn.Connect()

	user := src.User{Name: "Pinky Pie", Email: "test1@test.com", Password: "123456", Role: "Pony"}
	conn.CreateUser(&user)
}

func TestCall(t *testing.T) {
	println(src.Call())
}
