package middleware

import (
	"context"
	"net/http"

	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/27-sessions-standard/internal/sessionstore"
)

// Define a custom type for context keys to avoid collisions with other packages
type contextKey string

// Constant key for storing the username in request context
const userContextKey contextKey = "user"

// RequireSession is middleware that validates the session_token cookie.
// If valid, it attaches the username to the request context.
func RequireSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Look for session_token cookie
		cookie, err := r.Cookie("session_token")
		if err != nil || cookie.Value == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Check session store for matching user
		username, ok := sessionstore.GetSession(cookie.Value)
		if !ok {
			http.Error(w, "Session expired", http.StatusUnauthorized)
			return
		}

		// Store username in the request context using a safe custom key
		ctx := context.WithValue(r.Context(), userContextKey, username)

		// Pass the updated request context to the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserFromContext extracts the username string from request context.
func GetUserFromContext(r *http.Request) (string, bool) {
	username, ok := r.Context().Value(userContextKey).(string)
	return username, ok
}

/*
🧠 SAFELY ENFORCING SESSION AUTH — CONTEXT KEY BEST PRACTICES

✅ What Happens Here:
- We wrap protected routes with `RequireSession` middleware to ensure a valid session exists.
- If a valid session is found, we store the username in the context using a **custom type**.
- The new `GetUserFromContext()` helper makes it easy to retrieve the user later in handlers.

✅ Why This Matters:
- Using a custom context key type prevents accidental overwrites or conflicts across packages.
- Middleware enforces access control while keeping route handlers clean and focused.
- Separating context access into a helper improves reuse and testability.

✅ Key Concepts:
| Concept                       | Purpose                                                           |
|-------------------------------|-------------------------------------------------------------------|
| `type contextKey string`      | Creates a custom type to safely use as a context map key          |
| `context.WithValue()`         | Attaches data to a request lifecycle for downstream access        |
| `r.Context().Value(...)`      | Retrieves contextual values like "who is logged in?"              |
| Middleware + context pairing  | Ideal for passing auth, locale, or tracing data in clean apps     |

🔐 Bonus:
- You’ve now implemented **secure, idiomatic Go middleware** with production-ready patterns.
- This approach scales to RBAC, request tracing, or tenant-aware routing with minimal change.

📚 Up Next:
- Add role-based access logic
- Integrate a session timeout/refresh mechanism
- Store sessions in Redis or PostgreSQL for persistence
*/
