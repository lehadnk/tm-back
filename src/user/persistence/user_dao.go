package persistence

import (
	_ "github.com/lib/pq"
	"log"
	"tm/src/common"
	"tm/src/user/dto"
)

type UserDao struct {
	common.AbstractDao
}

func NewUserDao() *UserDao {
	var newUserDao = UserDao{}
	newUserDao.Connect()
	return &newUserDao
}

func (dbc *UserDao) CreateUser(user *dto.User) {
	log.Println("Creating user: " + user.Email)
	var userId int
	err := dbc.Db.QueryRow(
		"INSERT INTO users(name, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id",
		user.Name, user.Email, user.Password, user.Role).Scan(&userId)
	if err != nil {
		log.Println("Could not create user: " + err.Error())
	}
	user.Id = userId
}

func (dbc *UserDao) GetUserById(userId int) *dto.User {
	user := dto.User{}
	err := dbc.Db.Get(
		&user,
		"SELECT * from users WHERE id = $1", userId)
	if err != nil {
		log.Println("Could not find user: " + err.Error())
		return nil
	}

	return &user
}

func (dbc *UserDao) GetUserByEmailAndPassword(email string, password string) *dto.User {
	user := dto.User{}
	err := dbc.Db.Get(
		&user,
		"SELECT * from users WHERE email = $1 and password = $2", email, password)
	if err != nil {
		log.Println("Could not select user: " + err.Error())
	}
	return &user
}

func (dbc *UserDao) GetUsersList(sort string, page int, pageSize int) []dto.User {
	var users []dto.User
	var offset = (page - 1) * pageSize

	err := dbc.Db.Select(
		&users, "SELECT * from users ORDER BY $1 LIMIT $2 OFFSET $3", sort, pageSize, offset)
	if err != nil {
		log.Fatalln("Could not select users: " + err.Error())
	}
	return users
}

func (dbc *UserDao) GetUsersCount() int {
	var usersCount int
	err := dbc.Db.Get(
		&usersCount,
		"SELECT count(*) from users")
	if err != nil {
		log.Fatalln("Could not obtain user count: " + err.Error())
	}
	return usersCount
}

func (dbc *UserDao) EditUser(userId int, name string, email string, password string, role string) {
	if password != "" {
		_, err := dbc.Db.Exec(
			"UPDATE users SET name = $1, email = $2, password = $3, role = $4 WHERE id = $5",
			name, email, password, role, userId)
		if err != nil {
			log.Println("Could not update user: " + err.Error())
		}
	}

	_, err := dbc.Db.Exec(
		"UPDATE users SET name = $1, email = $2, role = $3 WHERE id = $4",
		name, email, role, userId)
	if err != nil {
		log.Println("Could not update user: " + err.Error())
	}
}

func (dbc *UserDao) DeleteUser(userId int) {
	_, err := dbc.Db.Exec("DELETE FROM users WHERE id = $1", userId)
	if err != nil {
		log.Println("Could not delete user: " + err.Error())
	}
}
