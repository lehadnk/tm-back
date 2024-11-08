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

func (dbc *UserDao) GetUserByEmailAndPassword(email string, password string) *domain.User {
	user := domain.User{}
	err := dbc.Db.Get(
		&user,
		"SELECT * from users WHERE email = $1 and password = $2", email, password)
	if err != nil {
		log.Fatalln(errors.New("could not get user"))
	}
	return &user
}

func (dbc *UserDao) GetUsersList(sort string, page int, pageSize int) []domain.User {
	var users []domain.User
	var offset = (page - 1) * pageSize

	err := dbc.Db.Select(
		&users, "SELECT * from users ORDER BY $1 LIMIT $2 OFFSET $3", sort, pageSize, offset)
	if err != nil {
		log.Fatalln(errors.New("could not get users"))
	}
	return users
}

func (dbc *UserDao) GetUsersCount() int {
	var usersCount int
	err := dbc.Db.Select(
		&usersCount,
		"SELECT count from users")
	if err != nil {
		log.Fatalln(errors.New("could not get count"))
	}
	return usersCount
}

func (dbc *UserDao) EditUser(name string, email string, password string, role string, userId int) {
	if password != "" {
		_, err := dbc.Db.Exec(
			"UPDATE users SET name = $1, email = $2, password = $3, role = $4 WHERE id = $5",
			name, email, password, role, userId)
		if err != nil {
			log.Fatalln(errors.New("could not update user"))
		}
	}

	_, err := dbc.Db.Exec(
		"UPDATE users SET name = $1, email = $2, role = $3 WHERE id = $4",
		name, email, role, userId)
	if err != nil {
		log.Fatalln(errors.New("could not update user"))
	}
}

func (dbc *UserDao) DeleteUser(userId int) {
	_, err := dbc.Db.Exec("DELETE FROM users WHERE id = $1", userId)
	if err != nil {
		log.Fatalln(errors.New("could not delete user"))
	}
}
