package repository

import (
	"database/sql"
	"errors"

	"taskflow-backend/internal/domain"
)

// ErrNotFound is returned when a record does not exist.
var ErrNotFound = errors.New("record not found")

type userRepository struct {
	db *sql.DB
}

// NewUserRepository creates a domain.UserRepository backed by SQLite.
func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(u *domain.User) error {
	_, err := r.db.Exec(
		`INSERT INTO users (id, name, email, password_hash, created_at, updated_at)
		 VALUES (?, ?, ?, ?, ?, ?)`,
		u.ID, u.Name, u.Email, u.PasswordHash, u.CreatedAt, u.UpdatedAt,
	)
	return err
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	row := r.db.QueryRow(
		`SELECT id, name, email, password_hash, created_at, updated_at
		 FROM users WHERE email = ?`, email,
	)
	return scanUser(row)
}

func (r *userRepository) FindByID(id string) (*domain.User, error) {
	row := r.db.QueryRow(
		`SELECT id, name, email, password_hash, created_at, updated_at
		 FROM users WHERE id = ?`, id,
	)
	return scanUser(row)
}

func scanUser(row *sql.Row) (*domain.User, error) {
	u := &domain.User{}
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.CreatedAt, &u.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return u, nil
}
