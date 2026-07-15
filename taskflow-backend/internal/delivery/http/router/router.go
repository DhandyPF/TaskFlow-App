package router

import (
	"net/http"

	"taskflow-backend/internal/delivery/http/handler"
	"taskflow-backend/internal/delivery/http/middleware"
)

// New builds the full HTTP router for the TaskFlow API.
func New(authHandler *handler.AuthHandler, todoHandler *handler.TodoHandler, jwtSecret string) http.Handler {
	mux := http.NewServeMux()

	// Health check
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	// Public auth endpoints
	mux.HandleFunc("POST /api/v1/auth/register", authHandler.Register)
	mux.HandleFunc("POST /api/v1/auth/login", authHandler.Login)

	// Protected todo endpoints
	protected := http.NewServeMux()
	protected.HandleFunc("GET /api/v1/todos", todoHandler.List)
	protected.HandleFunc("POST /api/v1/todos", todoHandler.Create)
	protected.HandleFunc("PUT /api/v1/todos/{id}", todoHandler.Update)
	protected.HandleFunc("DELETE /api/v1/todos/{id}", todoHandler.Delete)

	authMiddleware := middleware.Auth(jwtSecret)
	mux.Handle("/api/v1/todos", authMiddleware(protected))
	mux.Handle("/api/v1/todos/", authMiddleware(protected))

	return middleware.CORS(middleware.Logger(mux))
}
