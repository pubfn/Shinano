package repository

import (
	"database/sql"

	"Shinano/db"
)

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
