package db

import (
	"database/sql"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(db *sql.DB) error {
	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"sqlite",
		driver,
	)
	if err != nil {
		return err
	}

	err = m.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	}

	return err
}
