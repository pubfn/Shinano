package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func Init() *sql.DB {
	db, err := sql.Open("sqlite", "./main.db")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	if err := Migrate(db); err != nil {
		log.Fatal(err)
	}
	return db
}

func NewCommader(accId int, db *sql.DB) {
	_, err := db.Exec("INSERT INTO commander (id, name, level, exp) VALUES (?, ?, ?, ?)", accId, "7sdu", 6, 0)
	if err != nil {
		fmt.Println("1")
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO resourcelist (type, num) VALUES (?, ?)", 4, 5252)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO resourcelist (type, num) VALUES (?, ?)", 1, 6767)
	if err != nil {
		fmt.Println("3")
		log.Fatal(err)
	}
}
