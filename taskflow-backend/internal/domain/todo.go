package domain

import "time"

// Priority enumerates allowed todo priority levels.
type Priority string

const (
	PriorityLow    Priority = "LOW"
	PriorityMedium Priority = "MEDIUM"
	PriorityHigh   Priority = "HIGH"
)

// Todo represents a single task owned by a user.
type Todo struct {
	ID          string     `json:"id"`
	UserID      string     `json:"-"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	IsCompleted bool       `json:"is_completed"`
	Priority    Priority   `json:"priority"`
	DueDate     *time.Time `json:"due_date"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// CreateTodoRequest is the payload to create a task.
type CreateTodoRequest struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Priority    Priority   `json:"priority"`
	DueDate     *time.Time `json:"due_date"`
}

// UpdateTodoRequest is the payload to update a task.
type UpdateTodoRequest struct {
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	IsCompleted *bool      `json:"is_completed"`
	Priority    *Priority  `json:"priority"`
	DueDate     *time.Time `json:"due_date"`
}

// TodoFilter carries optional query parameters for listing tasks.
type TodoFilter struct {
	IsCompleted *bool
	SortBy      string // "due_date" or "created_at"
}

// TodoRepository defines persistence operations for todos.
type TodoRepository interface {
	Create(t *Todo) error
	FindAllByUser(userID string, filter TodoFilter) ([]*Todo, error)
	FindByID(id string) (*Todo, error)
	Update(t *Todo) error
	Delete(id string) error
}
