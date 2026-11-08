package db

import (
	"database/sql"
	"fmt"
	"log"

	"Shinano/cfg"

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

func NewCommader(accId int, db *sql.DB) {
	_, err := db.Exec("INSERT INTO commander (id, name, level, exp) VALUES (?, ?, ?, ?)", accId, "7sdu", 6, 0)
	_, err = db.Exec("INSERT INTO resourcelist (type, num) VALUES (?, ?)", 4, 5252)
	_, err = db.Exec("INSERT INTO resourcelist (type, num) VALUES (?, ?)", 1, 6767)
	_, err = db.Exec("insert into shiplist (id, templateId, level,exp,islocked,skinid) values (?,?,?,?,?,?)", 1, 307084, 67, 100, 0, 0)

	for _, v := range cfg.Skillid[307084].Id {
		_, err = db.Exec("insert into skillidlist (skillid, skilllvl, skillexp, shipid) values (?,?,?,?)", v, 2, 50, 307084)
		fmt.Println(v)
	}

	if err != nil {
		log.Fatal(err)
	}
}
