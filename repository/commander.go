package repository

import (
	"database/sql"

	"Shinano/db"
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
		err = db.NewCommader(int(accountID), b)
	}
	return nil
}
