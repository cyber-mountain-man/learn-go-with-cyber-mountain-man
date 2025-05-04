package handlers

import (
	"net/http"

	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/27-sessions-standard/middleware"
	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/27-sessions-standard/internal/sessionstore"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the homepage. Visit /login to authenticate."))
}

func Login(w http.ResponseWriter, r *http.Request) {
	sessionID := sessionstore.CreateSession("demo_user")

	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, cookie)

	w.Write([]byte("✅ Logged in. Visit /dashboard"))
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err == nil {
		sessionstore.DeleteSession(cookie.Value)
	}

	expired := &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	}
	http.SetCookie(w, expired)

	w.Write([]byte("🚪 Logged out."))
}

// ✅ Updated to safely access user via custom context key
func Dashboard(w http.ResponseWriter, r *http.Request) {
	username, ok := middleware.GetUserFromContext(r)
	if !ok {
		http.Error(w, "Could not extract user from context", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("🔐 Welcome to your dashboard, " + username))
}

/*
🧠 SESSION-AWARE HANDLERS — GO STANDARD LIBRARY EDITION

✅ What Happens Here:
- `Home` serves a public message with a prompt to authenticate.
- `Login` generates a random session ID for a hardcoded user, stores it server-side, and sets a secure cookie.
- `Logout` invalidates the session both server-side and client-side by deleting the token and expiring the cookie.
- `Dashboard` is a protected route that reads the username from request context (populated by middleware).

✅ Why This Matters:
- Demonstrates how to build a minimal authentication workflow using only Go’s standard library.
- Shows the correct use of middleware + context to pass session data across request handlers.
- Provides a real-world foundation for role-based access, session timeout, and secure routing.

✅ Key Concepts:
| Handler     | Responsibility                                      |
|-------------|------------------------------------------------------|
| `Home`      | Public-facing content, no session logic             |
| `Login`     | Creates a session token and returns it as a cookie  |
| `Logout`    | Deletes the session token and clears the cookie     |
| `Dashboard` | Reads user identity from context (set by middleware) |

🔐 Security Highlights:
- All session cookies are `HttpOnly` and use `SameSite=Strict` by default.
- Session IDs are generated using `crypto/rand` in the `sessionstore` package.
- The context key uses a custom type to prevent collisions (Go best practice).

📚 Up Next:
- Replace hardcoded users with real login forms and credential validation.
- Implement session timeout logic or sliding expiration.
- Introduce persistent session storage (e.g., Redis, PostgreSQL).
*/

