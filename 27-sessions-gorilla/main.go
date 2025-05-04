package main

import (
	"log"
	"net/http"

	// Import route handlers and middleware from local packages
	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/27-sessions-gorilla/handlers"
	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/27-sessions-gorilla/middleware"
)

func main() {
	// Create a new HTTP multiplexer (router)
	mux := http.NewServeMux()

	// Public routes
	mux.HandleFunc("/", handlers.Home)     // Accessible to everyone
	mux.HandleFunc("/login", handlers.Login)   // Creates session and sets cookie
	mux.HandleFunc("/logout", handlers.Logout) // Destroys session and clears cookie

	// Protected route — requires session token to access
	mux.Handle("/dashboard", middleware.RequireSession(http.HandlerFunc(handlers.Dashboard)))

	// Start the HTTP server on port 8080
	log.Println("🍪 Gorilla session server running at http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}

/*
🧠 GO WEB SERVER — GORILLA SESSIONS EDITION

✅ What Happens Here:
- This entry point registers all HTTP routes using `http.NewServeMux`.
- It introduces protected routing via Gorilla sessions — session tokens are securely stored and validated.
- The `/dashboard` route uses middleware to enforce authentication based on session presence.
- The server listens for incoming HTTP requests on `localhost:8080`.

✅ Why This Matters:
- This is a **real-world pattern**: clearly separated concerns (handlers vs middleware), cookie-based auth, modular routing.
- Using `gorilla/sessions` enables secure, persistent session management using encrypted cookies or server-side storage.
- Middleware abstracts away session validation so that handlers focus only on business logic.

✅ Key Concepts:
| Concept                      | Purpose                                                  |
|------------------------------|----------------------------------------------------------|
| `gorilla/sessions`           | Secure cookie-based session management                  |
| `http.HandlerFunc`           | Converts a function into a handler                      |
| `middleware.RequireSession`  | Validates the session and injects context               |
| Modular routing              | Cleanly separates public vs protected routes            |

🔐 Security Considerations:
- Gorilla sessions support encryption and authentication of cookies.
- Cookie options such as `HttpOnly`, `Secure`, and `SameSite` can be configured centrally.

📚 Up Next:
- Implement role-based access control
- Store sessions in Redis or filesystem for persistence
- Add flash messages or user profiles tied to sessions
*/
