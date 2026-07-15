package usecase

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"taskflow-backend/internal/domain"
	"taskflow-backend/internal/repository"
)

// ErrForbidden is returned when a user tries to act on another user's task.
var ErrForbidden = errors.New("you do not have access to this task")

// ErrNotFound is returned when a task does not exist.
var ErrNotFound = errors.New("task not found")

// TodoUsecase implements task management business logic.
type TodoUsecase struct {
	todos domain.TodoRepository
}

// NewTodoUsecase builds a TodoUsecase.
func NewTodoUsecase(todos domain.TodoRepository) *TodoUsecase {
	return &TodoUsecase{todos: todos}
}

// Create adds a new task owned by userID.
func (u *TodoUsecase) Create(userID string, req domain.CreateTodoRequest) (*domain.Todo, error) {
	priority := req.Priority
	if priority == "" {
		priority = domain.PriorityMedium
	}

	now := time.Now().UTC()
	todo := &domain.Todo{
		ID:          uuid.NewString(),
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		IsCompleted: false,
		Priority:    priority,
		DueDate:     req.DueDate,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := u.todos.Create(todo); err != nil {
		return nil, err
	}
	return todo, nil
}

// List returns every task owned by userID, honoring optional filters.
func (u *TodoUsecase) List(userID string, filter domain.TodoFilter) ([]*domain.Todo, error) {
	return u.todos.FindAllByUser(userID, filter)
}

// Update mutates fields of a task the caller owns.
func (u *TodoUsecase) Update(userID, id string, req domain.UpdateTodoRequest) (*domain.Todo, error) {
	todo, err := u.get(userID, id)
	if err != nil {
		return nil, err
	}

	if req.Title != nil {
		todo.Title = *req.Title
	}
	if req.Description != nil {
		todo.Description = *req.Description
	}
	if req.IsCompleted != nil {
		todo.IsCompleted = *req.IsCompleted
	}
	if req.Priority != nil {
		todo.Priority = *req.Priority
	}
	if req.DueDate != nil {
		todo.DueDate = req.DueDate
	}
	todo.UpdatedAt = time.Now().UTC()

	if err := u.todos.Update(todo); err != nil {
		return nil, err
	}
	return todo, nil
}

// Delete removes a task the caller owns.
func (u *TodoUsecase) Delete(userID, id string) error {
	if _, err := u.get(userID, id); err != nil {
		return err
	}
	return u.todos.Delete(id)
}

// get fetches a task and enforces ownership, preventing IDOR access.
func (u *TodoUsecase) get(userID, id string) (*domain.Todo, error) {
	todo, err := u.todos.FindByID(id)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	if todo.UserID != userID {
		return nil, ErrForbidden
	}
	return todo, nil
}
