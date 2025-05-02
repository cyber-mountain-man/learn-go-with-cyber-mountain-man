package main

import (
	"log"
	"net/http"

	// Chi is a lightweight, idiomatic router for building HTTP services in Go
	"github.com/go-chi/chi/v5"

	// Internal packages for DB connection and route handlers
	"learn-go-with-cyber-mountain-man/26-databases/postgres/internal/db"
	"learn-go-with-cyber-mountain-man/26-databases/postgres/internal/handlers"
)

func main() {
	// Establish database connection using the custom db package
	conn := db.Connect()
	// Ensure the DB connection is gracefully closed when main() exits
	defer conn.Close()

	// Create a new HTTP router using chi
	r := chi.NewRouter()

	// Define a group of routes under the /users path
	r.Route("/users", func(r chi.Router) {
		// GET /users        → List all users
		r.Get("/", handlers.ListUsers(conn))
		// POST /users       → Create a new user
		r.Post("/", handlers.CreateUser(conn))
		// GET /users/{id}   → Fetch a specific user by ID
		r.Get("/{id}", handlers.GetUser(conn))
		// PUT /users/{id}   → Update user by ID
		r.Put("/{id}", handlers.UpdateUser(conn))
		// DELETE /users/{id}→ Delete user by ID
		r.Delete("/{id}", handlers.DeleteUser(conn))
	})

	// Log server startup message with port info
	log.Println("Server running on http://localhost:8080")

	// Start the HTTP server and bind it to port 8080
	// This blocks the main goroutine until the server stops
	http.ListenAndServe(":8080", r)
}

/*
🧠 GO WEB SERVER + ROUTER SETUP — MAIN ENTRY POINT (main.go)

✅ What Happens Here:
- This is the application's main entry point that initializes the web server and route configuration.
- It sets up a router using the Chi library and binds each CRUD operation to the `/users` route.
- The server listens on port 8080 and serves HTTP requests using idiomatic Go tools.

✅ Why This Matters:
- This structure reflects **production-ready Go practices** — separation of concerns, graceful shutdown, and clean routing.
- Chi is fast and composable, making it perfect for building RESTful APIs without heavy frameworks.
- Clean route definitions improve maintainability and developer onboarding.

✅ Key Concepts:
| Concept                   | Explanation |
|----------------------------|-------------|
| `chi.NewRouter()`          | Creates a new instance of a Chi router |
| `r.Route("/users", ...)`   | Group routes under a shared prefix |
| `handlers.*(conn)`         | Injects the shared DB connection into handlers |
| `http.ListenAndServe`      | Starts the server and blocks until it stops |
| `defer conn.Close()`       | Ensures DB resources are released cleanly on exit |

🔔 Bonus Tip:
- Chi supports middleware out of the box — you can easily add logging, auth, or rate limiting per route or globally.
- The router is **goroutine-safe** and the server automatically handles multiple concurrent HTTP requests.

📚 Up Next:
- Add middleware for logging or CORS
- Create versioned APIs (e.g., `/api/v1/users`)
- Write integration tests for your routes
*/
