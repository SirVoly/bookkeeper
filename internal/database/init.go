package database

import (
	"database/sql"
    _ "modernc.org/sqlite"
)

func InitDatabase() (*Queries, error) {
	db, err := sql.Open("sqlite", "bookkeeper.db")
	if err != nil {
		return nil, err
	}

	queries := New(db)
	return queries, nil
}
