package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"taskflow-backend/internal/domain"
	"taskflow-backend/internal/usecase"
)

// AuthHandler exposes HTTP endpoints for registration and login.
type AuthHandler struct {
	auth *usecase.AuthUsecase
}

// NewAuthHandler builds an AuthHandler.
func NewAuthHandler(auth *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{auth: auth}
}

// Register handles POST /api/v1/auth/register.
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req domain.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if strings.TrimSpace(req.Name) == "" || strings.TrimSpace(req.Email) == "" || len(req.Password) < 8 {
		writeError(w, http.StatusUnprocessableEntity, "name, email and a password of at least 8 characters are required")
		return
	}

	user, err := h.auth.Register(req)
	if errors.Is(err, usecase.ErrEmailTaken) {
		writeError(w, http.StatusConflict, "email already registered")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not register user")
		return
	}

	writeSuccess(w, http.StatusCreated, "User registered successfully", map[string]string{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}

// Login handles POST /api/v1/auth/login.
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req domain.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	token, _, err := h.auth.Login(req)
	if errors.Is(err, usecase.ErrInvalidCredentials) {
		writeError(w, http.StatusUnauthorized, "invalid email or password")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not process login")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "Login successful",
		"token":   token,
	})
}
