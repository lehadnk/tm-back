package common

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

type AbstractDao struct {
	Db *sqlx.DB
}

func (dbc *AbstractDao) Connect() {
	db_name := os.Getenv("TM_DB_NAME")
	if db_name == "" {
		db_name = "postgres"
	}

	host := os.Getenv("TM_DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("TM_DB_PORT")
	if port == "" {
		port = "5432"
	}

	user := os.Getenv("TM_DB_USER")
	if user == "" {
		user = "postgres"
	}

	password := os.Getenv("TM_DB_PASSWORD")
	if password == "" {
		password = "postgres"
	}

	ssl := os.Getenv("TM_DB_SSL")
	if ssl == "" {
		ssl = "disable"
	}

	datasource := "host=" + host + " port=" + port + " user=" + user + " dbname=" + db_name + " password=" + password + " sslmode=" + ssl

	db, err := sqlx.Connect("postgres", datasource)
	if err != nil {
		log.Fatalln("Could not connect to database: " + err.Error())
	}
	dbc.Db = db

	fmt.Println("Connected to postgres instance at " + host + ":" + port + " with db name " + db_name)
}
