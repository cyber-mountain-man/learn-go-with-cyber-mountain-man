package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"learn-go-with-cyber-mountain-man/26-databases/mssql/internal/db"
	"learn-go-with-cyber-mountain-man/26-databases/mssql/internal/handlers"
)

func main() {
	// Establish DB connection using MSSQL driver
	conn := db.Connect()
	defer conn.Close()

	// Initialize Chi router
	r := chi.NewRouter()

	// Register user-related routes
	r.Route("/users", func(r chi.Router) {
		r.Get("/", handlers.ListUsers(conn))     // GET all users
		r.Post("/", handlers.CreateUser(conn))   // POST new user
		r.Get("/{id}", handlers.GetUser(conn))   // GET single user
		r.Put("/{id}", handlers.UpdateUser(conn)) // PUT update user
		r.Delete("/{id}", handlers.DeleteUser(conn)) // DELETE user
	})

	// Start the HTTP server
	log.Println("✅ Server running on http://localhost:8080 (MSSQL)")
	http.ListenAndServe(":8080", r)
}
