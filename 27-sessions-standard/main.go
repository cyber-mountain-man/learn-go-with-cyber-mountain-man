package main

import (
	"log"
	"net/http"

	// Importing our handlers and middleware packages
	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/27-sessions-standard/handlers"
	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/27-sessions-standard/middleware"
)

func main() {
	// Create a new HTTP request multiplexer (router)
	mux := http.NewServeMux()

	// 🔓 Public endpoints
	mux.HandleFunc("/", handlers.Home)       // Accessible by anyone
	mux.HandleFunc("/login", handlers.Login) // Simulates login and sets session cookie
	mux.HandleFunc("/logout", handlers.Logout) // Clears the session

	// 🔐 Protected endpoint wrapped with session-checking middleware
	mux.Handle("/dashboard", middleware.RequireSession(http.HandlerFunc(handlers.Dashboard)))
	// If the session is valid, the request proceeds to Dashboard handler
	// If not, it returns 401 Unauthorized

	// Start the web server on localhost:8080
	log.Println("🔐 Server running at http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}

/*
🧠 SESSION-AWARE SERVER — LESSON 27 OVERVIEW

✅ What Happens Here:
- This file wires up all application routes using the standard library.
- Public endpoints like `/`, `/login`, and `/logout` are freely accessible.
- The `/dashboard` route is protected by middleware that checks for a valid session cookie.
- If a user is not authenticated (no valid cookie), they're denied access.

✅ Why This Matters:
- This pattern is the **core of modern authentication** in web applications.
- Even without external packages, Go enables secure session management via cookies and middleware.
- Helps build an understanding of how session-based auth works from scratch.

✅ Key Concepts:
| Concept                      | Purpose                                                  |
|------------------------------|----------------------------------------------------------|
| `http.NewServeMux()`         | Lightweight built-in router                              |
| `HandleFunc` / `Handle`      | Maps paths to handlers or middleware                     |
| Middleware                   | Intercepts HTTP requests to enforce security             |
| Cookie-based session         | Tracks user identity without full login systems          |
| Context propagation          | Safely injects user identity into handlers               |

🔐 Security Tip:
- Cookies are marked `HttpOnly` and `SameSite` to mitigate XSS and CSRF risks.
- Each session token is securely generated with 32 bytes of random data.

📚 Up Next:
- Improve sessions using secure stores like `gorilla/sessions`
- Add login forms and password hashing
- Store session tokens in Redis or a database for scalability
*/