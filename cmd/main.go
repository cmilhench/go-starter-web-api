package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cmilhench/go-starter-web-api/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.Home)
	mux.HandleFunc("GET /api/health", handlers.Health)
	mux.HandleFunc("GET /api/users/{id}", handlers.GetUser)
	mux.HandleFunc("GET /api/users", handlers.GetUsers)
	mux.HandleFunc("POST /api/users", handlers.CreateUser)
	mux.HandleFunc("PUT /api/users/{id}", handlers.UpdateUser)
	mux.HandleFunc("DELETE /api/users/{id}", handlers.DeleteUser)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Server starting on port %s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
