package persistence

import (
	"awesomeProject/src/jwt/domain"
	"errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type DbConnection struct {
	db *sqlx.DB
}

func (dbc *DbConnection) Connect() {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		log.Fatalln(errors.New("could not connect to database"))
	}
	dbc.db = db
}

func (dbc *DbConnection) CreateUser(user *domain.User) {
	var userId int
	err := dbc.db.QueryRow(
		"INSERT INTO users(name, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id",
		user.Name, user.Email, user.Password, user.Role).Scan(&userId)
	if err != nil {
		log.Fatalln(errors.New("could not create user"))
	}
}
