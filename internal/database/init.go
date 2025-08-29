package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func InitQueries() (*Queries, error) {
	db, err := InitDatabase()
	if err != nil {
		return nil, err
	}

	queries := New(db)
	return queries, nil
}

func InitDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "bookkeeper.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}