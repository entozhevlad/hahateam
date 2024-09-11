package sqlite

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.NewStorage"
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS USERS(
	    login TEXT NOT NULL PRIMARY KEY,
	    password TEXT NOT NULL UNIQUE,
	    company TEXT NOT NULL);
	CREATE INDEX IF NOT EXISTS idx_login ON USERS(login);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}
