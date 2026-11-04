package constructor

import "database/sql"

type Builder struct {
	DB *sql.DB
}

func NewBuilder(db *sql.DB) *Builder {
	return &Builder{
		DB: db,
	}
}
