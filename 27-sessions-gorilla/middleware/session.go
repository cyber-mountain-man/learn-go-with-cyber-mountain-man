package middleware

import (
	"net/http"

	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/27-sessions-gorilla/internal/config"
)

// RequireSession wraps a protected route and checks if a valid session exists.
func RequireSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := config.Store.Get(r, "session")

		// Check for session value "username"
		if username, ok := session.Values["username"].(string); !ok || username == "" {
			http.Error(w, "🔒 Unauthorized. Please login first.", http.StatusUnauthorized)
			return
		}

		// Session is valid; proceed with the request
		next.ServeHTTP(w, r)
	})
}

/*
🧠 SESSION MIDDLEWARE — GORILLA WRAPPER

✅ What Happens Here:
- This middleware checks whether the incoming HTTP request contains a valid session.
- Specifically, it looks for the key `username` in the session — if missing, access is denied.
- If the session is valid, it allows the request to proceed to the next handler.

✅ Why This Matters:
- Centralizes access control logic: write it once, apply it to many routes.
- Keeps your business logic clean and free of repetitive session checks.
- Encourages separation of concerns: auth belongs in middleware, not in views.

✅ Key Concepts:
| Concept                    | Purpose                                  |
|----------------------------|------------------------------------------|
| `Store.Get(r, "session")`  | Loads the session associated with request |
| `session.Values[...]`      | Reads or writes session-level data        |
| `http.HandlerFunc`         | Adapter to treat functions as middleware  |

🔐 Best Practices:
- Always validate session data before using it
- Consider adding logging or redirect logic for failed auth

📚 Up Next:
- Add secure configuration for the session store
- Build persistent storage for sessions using Redis or database
*/

