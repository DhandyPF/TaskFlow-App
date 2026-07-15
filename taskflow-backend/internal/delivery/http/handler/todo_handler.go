package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"taskflow-backend/internal/delivery/http/middleware"
	"taskflow-backend/internal/domain"
	"taskflow-backend/internal/usecase"
)

// TodoHandler exposes HTTP endpoints for task management.
type TodoHandler struct {
	todos *usecase.TodoUsecase
}

// NewTodoHandler builds a TodoHandler.
func NewTodoHandler(todos *usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{todos: todos}
}

// List handles GET /api/v1/todos.
func (h *TodoHandler) List(w http.ResponseWriter, r *http.Request) {
	userID := middleware.UserIDFromContext(r.Context())

	filter := domain.TodoFilter{SortBy: r.URL.Query().Get("sort_by")}
	if raw := r.URL.Query().Get("is_completed"); raw != "" {
		if v, err := strconv.ParseBool(raw); err == nil {
			filter.IsCompleted = &v
		}
	}

	todos, err := h.todos.List(userID, filter)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not fetch tasks")
		return
	}
	if todos == nil {
		todos = []*domain.Todo{}
	}

	writeSuccess(w, http.StatusOK, "", todos)
}

// Create handles POST /api/v1/todos.
func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID := middleware.UserIDFromContext(r.Context())

	var req domain.CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if req.Title == "" {
		writeError(w, http.StatusUnprocessableEntity, "title is required")
		return
	}

	todo, err := h.todos.Create(userID, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not create task")
		return
	}

	writeSuccess(w, http.StatusCreated, "Task created successfully", todo)
}

// Update handles PUT /api/v1/todos/:id.
func (h *TodoHandler) Update(w http.ResponseWriter, r *http.Request) {
	userID := middleware.UserIDFromContext(r.Context())
	id := r.PathValue("id")

	var req domain.UpdateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	todo, err := h.todos.Update(userID, id, req)
	switch {
	case errors.Is(err, usecase.ErrNotFound):
		writeError(w, http.StatusNotFound, "task not found")
		return
	case errors.Is(err, usecase.ErrForbidden):
		writeError(w, http.StatusForbidden, "you do not have access to this task")
		return
	case err != nil:
		writeError(w, http.StatusInternalServerError, "could not update task")
		return
	}

	writeSuccess(w, http.StatusOK, "Task updated successfully", todo)
}

// Delete handles DELETE /api/v1/todos/:id.
func (h *TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userID := middleware.UserIDFromContext(r.Context())
	id := r.PathValue("id")

	err := h.todos.Delete(userID, id)
	switch {
	case errors.Is(err, usecase.ErrNotFound):
		writeError(w, http.StatusNotFound, "task not found")
		return
	case errors.Is(err, usecase.ErrForbidden):
		writeError(w, http.StatusForbidden, "you do not have access to this task")
		return
	case err != nil:
		writeError(w, http.StatusInternalServerError, "could not delete task")
		return
	}

	writeSuccess(w, http.StatusOK, "Task deleted successfully", nil)
}
