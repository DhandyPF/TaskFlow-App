package main

import (
	"log"
	"net/http"

	"taskflow-backend/internal/config"
	"taskflow-backend/internal/delivery/http/handler"
	"taskflow-backend/internal/delivery/http/router"
	"taskflow-backend/internal/repository"
	"taskflow-backend/internal/usecase"
)

func main() {
	cfg := config.Load()

	db, err := repository.NewDB(cfg.DatabaseDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	todoRepo := repository.NewTodoRepository(db)

	authUsecase := usecase.NewAuthUsecase(userRepo, cfg.JWTSecret)
	todoUsecase := usecase.NewTodoUsecase(todoRepo)

	authHandler := handler.NewAuthHandler(authUsecase)
	todoHandler := handler.NewTodoHandler(todoUsecase)

	r := router.New(authHandler, todoHandler, cfg.JWTSecret)

	log.Printf("TaskFlow API listening on :%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
