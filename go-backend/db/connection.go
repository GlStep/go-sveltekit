package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func InitDB(connectionURL string) {
	var err error
	db, err = sqlx.Open("postgres", connectionURL)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Error pinging database: ", err)
	}
}

func GetDB() *sqlx.DB {
	return db
}
