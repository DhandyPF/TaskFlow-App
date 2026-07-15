package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// NewDB opens the SQLite database and ensures the schema exists.
func NewDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn+"?_foreign_keys=on")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	if err := migrate(db); err != nil {
		return nil, err
	}
	return db, nil
}

func migrate(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL
	);

	CREATE TABLE IF NOT EXISTS todos (
		id TEXT PRIMARY KEY,
		user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		title TEXT NOT NULL,
		description TEXT,
		is_completed BOOLEAN NOT NULL DEFAULT 0,
		priority TEXT NOT NULL DEFAULT 'MEDIUM',
		due_date DATETIME,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL
	);

	CREATE INDEX IF NOT EXISTS idx_todos_user_id ON todos(user_id);
	`
	_, err := db.Exec(schema)
	return err
}
