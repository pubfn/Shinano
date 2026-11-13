package db

import (
	"database/sql"
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

func NewCommader(accId int, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	funcs := []func() error{
		func() error {
			_, err = tx.Exec("INSERT INTO commander (id, name, level, exp) VALUES (?, ?, ?, ?)", accId, "7sdu", 6, 0)
			return err
		},
		func() error {
			_, err = tx.Exec("INSERT INTO resourcelist (type, num) VALUES (?, ?)", 4, 5252)
			return err
		},
		func() error {
			_, err = tx.Exec("INSERT INTO resourcelist (type, num) VALUES (?, ?)", 1, 6767)
			return err
		},
		func() error {
			_, err = tx.Exec("insert into shiplist (id, templateId, level,exp,islocked,skinid) values (?,?,?,?,?,?)", 1, 307084, 67, 100, 0, 0)
			return err
		},
		func() error {
			_, err = tx.Exec("insert into inventory(id, count) values (?,?)", 20001, 6742)
			return err
		},
		func() error {
			for _, v := range cfg.SkillID[307084].ID {
				_, err = tx.Exec("insert into skillidlist (skillid, skilllvl, skillexp, shipid) values (?,?,?,?)", v, 2, 50, 307084)
			}
			return err
		},
	}

	for _, transaction := range funcs {
		if err := transaction(); err != nil {
			return err
		}
	}

	return tx.Commit()
}
