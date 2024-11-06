package persistence

import (
	"awesomeProject/src/common"
	"awesomeProject/src/user/domain"
	"errors"
	_ "github.com/lib/pq"
	"log"
)

type UserDao struct {
	common.AbstractDao
}

func NewUserDao() *UserDao {
	var newUserDao = UserDao{}
	newUserDao.Connect()
	return &newUserDao
}

func (dbc *UserDao) CreateUser(user *domain.User) {
	var userId int
	err := dbc.Db.QueryRow(
		"INSERT INTO users(name, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id",
		user.Name, user.Email, user.Password, user.Role).Scan(&userId)
	if err != nil {
		log.Fatalln(errors.New("could not create user"))
	}
	user.Id = userId
}

func (dbc *UserDao) GetUserById(userId int) *domain.User {
	user := domain.User{}
	err := dbc.Db.Get(
		&user,
		"SELECT * from users WHERE id = $1", userId)
	if err != nil {
		return nil
	}
	return &user
}

func (dbc *UserDao) DeleteUser(userId int) {
	_, err := dbc.Db.Exec("DELETE FROM users WHERE id = $1", userId)
	if err != nil {
		log.Fatalln(errors.New("could not delete user"))
	}
}
