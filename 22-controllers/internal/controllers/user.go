package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// UserController is a struct used to group related user route logic.
// In a full application, this is where you'd inject services, databases, etc.
type UserController struct {
	// Example: db *sql.DB
}

// NewUserController returns a new instance of UserController.
// Used in router setup to bind handler methods to specific route paths.
func NewUserController() *UserController {
	return &UserController{}
}

// GetUser handles GET /users/{userID}.
// It extracts the userID from the URL path, validates it, and responds accordingly.
func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	// Extract dynamic path variable {userID}
	userID := chi.URLParam(r, "userID")

	// Validate that it's a numeric ID
	if _, err := strconv.Atoi(userID); err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Mock response (would typically query DB here)
	fmt.Fprintf(w, "👤 User Profile: %s\n", userID)
}

// ListUsers handles GET /users.
// It returns a simple mock list of users (no DB logic yet).
func (uc *UserController) ListUsers(w http.ResponseWriter, r *http.Request) {
	// Mock user list response
	fmt.Fprintln(w, "📋 List of all users (mock data)")
}

/*
🧠 LESSON 22 - CONTROLLERS: LOGIC STRUCTURE WITH METHODS

✅ What You Learn:
- How to build a controller using Go structs and method receivers
- How to bind methods to routes for organized, scalable logic
- How to use dynamic parameters and validate them (e.g., userID)
- How to cleanly separate concerns: routing, logic, response

✅ Why This Matters:
- Enables dependency injection (DBs, services, loggers) in the future
- Follows scalable architecture seen in large backend services
- Keeps code testable, modular, and idiomatic to Go

This controller pattern makes it easy to extend your app with new features
(e.g., POST /users, PUT /users/{id}), while remaining clean and organized.
*/
