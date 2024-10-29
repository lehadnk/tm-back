package common

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"log"
)

type AbstractDao struct {
	Db *sqlx.DB
}

func (dbc *AbstractDao) Connect() {
	db, err := sqlx.Connect(
		"postgres",
		"host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		log.Fatalln(errors.New("could not connect to database"))
	}
	dbc.Db = db
}
