package db

import (
	"database/sql"
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

	db.Exec("PRAGMA foreign_keys = ON")

	if err := Migrate(db); err != nil {
		log.Fatal(err)
	}
	return db
}
