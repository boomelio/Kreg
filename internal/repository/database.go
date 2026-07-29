package repository

import (
	"database/sql"

	_"modernc.org/sqlite"
)

func Open(path string) (*sql.DB, error) {

	return sql.Open("sqlite", path)
}

func CreateSchema(db *sql.DB) error {
	const query = `
		CREATE TABLE IF NOT EXISTS monitors (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			url TEXT NOT NULL
		);
	`
	var _, err = db.Exec(query)
	return err
} 