package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/Akakazkz/go-task-manager-api/internal/handler"
	"github.com/Akakazkz/go-task-manager-api/internal/middleware"
	"github.com/Akakazkz/go-task-manager-api/internal/repository"
	"github.com/Akakazkz/go-task-manager-api/internal/service"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:7642004@localhost:5432/taskdb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	userRepo := repository.NewPostgresUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	taskRepo := repository.NewPostgresTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", handler.Health)
	mux.HandleFunc("/users", userHandler.Handle)
	mux.HandleFunc("/login", userHandler.Login)
	mux.Handle("/tasks", middleware.Auth(http.HandlerFunc(taskHandler.Handle)))
	mux.Handle("/tasks/", middleware.Auth(http.HandlerFunc(taskHandler.HandleByID)))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("starting server on :8080")
	log.Fatal(server.ListenAndServe())

}
