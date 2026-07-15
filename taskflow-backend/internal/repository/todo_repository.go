package repository

import (
	"database/sql"
	"errors"

	"taskflow-backend/internal/domain"
)

type todoRepository struct {
	db *sql.DB
}

// NewTodoRepository creates a domain.TodoRepository backed by SQLite.
func NewTodoRepository(db *sql.DB) domain.TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) Create(t *domain.Todo) error {
	_, err := r.db.Exec(
		`INSERT INTO todos (id, user_id, title, description, is_completed, priority, due_date, created_at, updated_at)
		 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		t.ID, t.UserID, t.Title, t.Description, t.IsCompleted, t.Priority, t.DueDate, t.CreatedAt, t.UpdatedAt,
	)
	return err
}

// FindAllByUser always scopes the query to user_id to prevent IDOR, per the
// PRD's non-functional security requirement.
func (r *todoRepository) FindAllByUser(userID string, filter domain.TodoFilter) ([]*domain.Todo, error) {
	query := `SELECT id, user_id, title, description, is_completed, priority, due_date, created_at, updated_at
	           FROM todos WHERE user_id = ?`
	args := []interface{}{userID}

	if filter.IsCompleted != nil {
		query += " AND is_completed = ?"
		args = append(args, *filter.IsCompleted)
	}

	switch filter.SortBy {
	case "due_date":
		query += " ORDER BY due_date ASC"
	default:
		query += " ORDER BY created_at DESC"
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*domain.Todo
	for rows.Next() {
		t := &domain.Todo{}
		if err := rows.Scan(&t.ID, &t.UserID, &t.Title, &t.Description, &t.IsCompleted,
			&t.Priority, &t.DueDate, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, rows.Err()
}

func (r *todoRepository) FindByID(id string) (*domain.Todo, error) {
	row := r.db.QueryRow(
		`SELECT id, user_id, title, description, is_completed, priority, due_date, created_at, updated_at
		 FROM todos WHERE id = ?`, id,
	)
	t := &domain.Todo{}
	err := row.Scan(&t.ID, &t.UserID, &t.Title, &t.Description, &t.IsCompleted,
		&t.Priority, &t.DueDate, &t.CreatedAt, &t.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *todoRepository) Update(t *domain.Todo) error {
	res, err := r.db.Exec(
		`UPDATE todos SET title = ?, description = ?, is_completed = ?, priority = ?, due_date = ?, updated_at = ?
		 WHERE id = ? AND user_id = ?`,
		t.Title, t.Description, t.IsCompleted, t.Priority, t.DueDate, t.UpdatedAt, t.ID, t.UserID,
	)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *todoRepository) Delete(id string) error {
	res, err := r.db.Exec(`DELETE FROM todos WHERE id = ?`, id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrNotFound
	}
	return nil
}
