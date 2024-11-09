package common

import (
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

type AbstractDao struct {
	Db *sqlx.DB
}

func (dbc *AbstractDao) Connect() {
	db_name := os.Getenv("T9TP1_DB_NAME")
	if db_name == "" {
		db_name = "postgres"
	}

	host := os.Getenv("T9TP1_DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("T9TP1_DB_PORT")
	if port == "" {
		port = "5432"
	}

	user := os.Getenv("T9TP1_DB_USER")
	if user == "" {
		user = "postgres"
	}

	password := os.Getenv("T9TP1_DB_PASSWORD")
	if password == "" {
		password = "postgres"
	}

	ssl := os.Getenv("T9TP1_DB_SSL")
	if ssl == "" {
		ssl = "disable"
	}

	datasource := "host=" + host + " port=" + port + " user=" + user + " dbname=" + db_name + " password=" + password + " sslmode=" + ssl

	db, err := sqlx.Connect("postgres", datasource)
	if err != nil {
		log.Fatalln("Could not connect to database: " + err.Error())
	}
	dbc.Db = db
}
