package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/27-sessions-gorilla/internal/config"
)

// Home displays the public landing page.
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("🏠 Welcome to the homepage. Go to /login to begin."))
}

// Login creates a new session and sets a secure cookie.
func Login(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, "session")
	session.Values["username"] = "demo_user"

	if err := session.Save(r, w); err != nil {
		log.Printf("❌ Failed to save session: %v", err)
		http.Error(w, "Could not save session", http.StatusInternalServerError)
		return
	}

	// ✅ Log raw Set-Cookie header for debugging
	for _, cookie := range w.Header()["Set-Cookie"] {
		log.Printf("🔍 Set-Cookie: %s", cookie)
	}

	w.Write([]byte("✅ Logged in with Gorilla sessions. Visit /dashboard"))
}


// Logout clears the session and expires the cookie.
func Logout(w http.ResponseWriter, r *http.Request) {
	// Retrieve session for current user
	session, err := config.Store.Get(r, "session")
	if err != nil {
		log.Printf("❌ Failed to get session for logout: %v", err)
		http.Error(w, "Could not retrieve session", http.StatusInternalServerError)
		return
	}

	// Mark session for deletion by setting MaxAge to -1
	session.Options.MaxAge = -1

	// Save the change to delete the cookie
	if err := session.Save(r, w); err != nil {
		log.Printf("❌ Failed to clear session: %v", err)
		http.Error(w, "Could not clear session", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("🚪 Logged out."))
}

// Dashboard is a protected route that requires a valid session.
func Dashboard(w http.ResponseWriter, r *http.Request) {
	// Attempt to retrieve session
	session, err := config.Store.Get(r, "session")
	if err != nil {
		log.Printf("❌ Failed to retrieve session: %v", err)
		http.Error(w, "Unauthorized access", http.StatusUnauthorized)
		return
	}

	// Safely extract the username from session data
	username, ok := session.Values["username"].(string)
	if !ok || username == "" {
		http.Error(w, "🔒 Unauthorized. Please login first.", http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("🔐 Welcome to your dashboard, %v", username)))
}

/*
🧠 HANDLERS — GORILLA SESSION EDITION (Final Version with Full Error Handling)

✅ What Happens Here:
- Each handler checks for session retrieval errors before accessing session data.
- `Login` sets a demo username and persists the session cookie.
- `Logout` marks the session for deletion and confirms removal with `session.Save()`.
- `Dashboard` ensures a valid, non-empty session exists before showing protected content.

✅ Why This Matters:
- Robust error checking helps prevent subtle bugs and improves security awareness.
- These handlers match real-world practices where session failures must be anticipated.

✅ Key Concepts:
| Concept                        | Purpose                                                  |
|-------------------------------|----------------------------------------------------------|
| `session.Save()`              | Persists session cookie updates (creation or deletion)   |
| Defensive programming         | Prevents undefined behavior when sessions fail           |
| `session.Values["key"].(T)`   | Safely casts session values with a type assertion        |
| Logging failures              | Supports debugging and runtime visibility                |

🔐 Best Practices:
- Never assume session data exists without validation.
- Clear sessions on logout to prevent reuse.
- Store only non-sensitive identifiers (not passwords, tokens) in the cookie.

📚 Up Next:
- Enforce session-based route protection with custom middleware
- Add user roles or permissions to session data
- Integrate Gorilla with Redis or PostgreSQL for scalable persistence
*/
