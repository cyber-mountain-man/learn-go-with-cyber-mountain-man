package routes

import (
	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/22-controllers/internal/controllers"
	"github.com/go-chi/chi/v5"
)

// SetupRouter initializes and returns a fully configured chi.Router
func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	// -----------------------------
	// 1️⃣ CONTROLLER INITIALIZATION
	// -----------------------------
	// Create an instance of the UserController.
	// In real apps, you might inject a database or service layer here.
	userController := controllers.NewUserController()

	// -----------------------------
	// 2️⃣ DEFINE ROUTES USING GROUPS
	// -----------------------------
	// All /users routes are handled here
	r.Route("/users", func(r chi.Router) {
		// GET /users/{userID} → handled by GetUser method
		r.Get("/{userID}", userController.GetUser)

		// GET /users → handled by ListUsers method
		r.Get("/", userController.ListUsers)
	})

	// Return the fully configured router to main.go
	return r
}

/*
🧠 LESSON 22 - CONTROLLERS: ROUTER LAYER

✅ What You Learn:
- How to cleanly wire route paths to controller methods
- How to use chi route groups for better organization (/users)
- How to initialize controller instances before binding routes

✅ Why This Matters:
- Keeps routing configuration separate from handler logic
- Makes it easy to grow your app modularly (e.g., add /products, /admin)
- Supports dependency injection — crucial for testing and real-world projects

This file bridges your entry point (main.go) and your business logic (controllers).
It’s where route paths meet functionality — cleanly and professionally.
*/
