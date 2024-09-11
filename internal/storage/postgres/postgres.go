package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(storagePath string) (*Storage, error) {
	const op = "storage.postgres.NewStorage"
	db, err := sql.Open("postgres", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS USERS(
	    id SERIAL PRIMARY KEY,
	    login TEXT NOT NULL UNIQUE,
	    password TEXT NOT NULL UNIQUE,
	    company TEXT NOT NULL UNIQUE);
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

// Register performs registration operation in the storage.
// It returns login if registration was successful, error otherwise.
func (s *Storage) Register(login string, password string, company string) (string, error) {
	const op = "storage.postgres.Register"
	// Execute the query against the database
	_, err := s.db.Exec("INSERT INTO USERS(login, password, company) VALUES($1, $2, $3)", login, password, company)
	if err != nil {
		// Return the error with a meaningful message
		return "", fmt.Errorf("%s: %w", op, err)
	}
	// Registration was successful, return the login
	return login, nil
}

// Login performs login operation in the storage.
// It returns company name if login was successful, error otherwise.
func (s *Storage) Login(login string, password string) (string, error) {
	const op = "storage.postgres.Login"
	var company string
	err := s.db.QueryRow("SELECT company FROM USERS WHERE login = $1 AND password = $2", login, password).Scan(&company)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return company, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}
