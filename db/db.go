package db

import (
	"database/sql"
	"log"
)

func Init() *sql.DB {
	db, err := sql.Open("sqlite3", "./main.db")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func NewCommader(accId int, db *sql.DB) {
	_, err := db.Exec("INSERT INTO commander (id, name, level, exp) VALUES (? ? ? ?)", accId, "penis?", 2, 0)
	if err != nil {
		log.Fatal(err)
	}
}
