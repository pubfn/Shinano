package repository

import (
	"database/sql"

	"Shinano/cfg"
)

type Commander struct {
	Level uint32
	Exp   uint32
	Id    uint32
}

func GetCommanderData(b *sql.DB) (*Commander, error) {
	commander := Commander{}

	if err := b.QueryRow("select id, level, exp from commander").Scan(&commander.Id, &commander.Level, &commander.Exp); err != nil {
		return nil, err
	}

	return &commander, nil
}

func CheckCommander(b *sql.DB, accountID int64) error {
	var exists bool

	err := b.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM commander WHERE id = ?)",
		accountID,
	).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		err = NewCommader(int(accountID), b)
	}
	return nil
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
			_, err = tx.Exec("INSERT INTO resourcelist (type, num) VALUES (?, ?)", 1, 676767)
			return err
		},
		func() error {
			return AddShip(tx, 307084)
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
