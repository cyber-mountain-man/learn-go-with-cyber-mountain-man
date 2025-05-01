package middleware

import (
	"net/http"
)

// RequireAuth is a middleware function that simulates basic authentication.
//
// In real-world applications, this is where you would verify access tokens,
// session cookies, or API keys before allowing a user to access protected routes.
func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// -----------------------------
		// 1️⃣ EXTRACT HEADER
		// -----------------------------
		// Look for a custom HTTP header "X-Auth" in the request.
		// This mimics how APIs receive tokens or secrets from clients.
		authHeader := r.Header.Get("X-Auth")

		// -----------------------------
		// 2️⃣ CHECK AUTH VALUE
		// -----------------------------
		// We're using a hardcoded secret value here for demonstration.
		// In a real app, you'd verify against a user database or token service.
		if authHeader != "secret123" {
			// If the header is missing or doesn't match, block access.
			// Return a 401 Unauthorized status with a friendly message.
			http.Error(w, "Unauthorized - missing or invalid X-Auth header", http.StatusUnauthorized)
			return
		}

		// -----------------------------
		// 3️⃣ PROCEED IF AUTHORIZED
		// -----------------------------
		// The request passes validation, so call the next handler.
		// This allows the route logic to proceed normally.
		next.ServeHTTP(w, r)
	})
}

/*
🧠 LESSON 23 - AUTHENTICATION MIDDLEWARE (SIMULATED)

✅ What This Teaches:
- Middleware in Go acts as a **filter** between the client and your route logic.
- `RequireAuth` demonstrates how to block or allow access **before** hitting route handlers.
- You learn how to:
  - Read request headers using `r.Header.Get(...)`
  - Return early with `http.Error()` and status codes like 401 Unauthorized
  - Compose middleware cleanly using `next.ServeHTTP(w, r)`

🔐 Why This Pattern Matters:
- This is the **first line of defense** in any secure backend application.
- It supports scalable, layered security: you can apply this only to sensitive endpoints like `/admin` or `/private`.
- Easy to plug in real authentication systems later (e.g., JWT, OAuth2, session cookies).

🛠️ Common Real-World Extensions:
- Check API keys or JWT tokens from headers
- Validate session IDs against a database
- Enforce HTTPS-only access
- Add rate-limiting or IP restrictions

📌 Best Practices:
- Don’t bloat route logic with access checks — keep them in middleware
- Make middleware composable and testable on its own
- Return **clear, consistent error messages** for debugging and API users

This is your gateway to building **zero-trust, security-first Go web services**.
It's small in code, but massive in real-world impact.
*/
