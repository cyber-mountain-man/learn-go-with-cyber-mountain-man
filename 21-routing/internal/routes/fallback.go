package routes

import (
	"fmt"
	"net/http"
)

// FallbackHandler handles requests to undefined routes
func FallbackHandler(w http.ResponseWriter, r *http.Request) {
	// Set HTTP status to 404 Not Found
	w.WriteHeader(http.StatusNotFound)

	// Return a custom message to the user
	fmt.Fprintln(w, "🚫 404 - Page Not Found")
}

/*
🧠 LESSON 21 - ROUTING: CUSTOM 404 FALLBACK HANDLER

✅ What You Learn:
- How to create a universal fallback for routes not matched by your router
- How to manually set the HTTP response status to 404
- How to serve a friendly and branded "Page Not Found" message

✅ Why This Matters:
- A custom fallback improves user experience compared to default server 404s
- Ensures unknown paths don't accidentally leak server internals
- Gives you full control over error responses — critical for APIs and production services
- Aligns with best practices for clean, complete, user-friendly server behavior

Setting up a good fallback handler is a small but essential step toward building reliable, polished web applications.

*/
