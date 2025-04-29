package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// UserRoutes registers the /users routes on the provided chi.Router
func UserRoutes(r chi.Router) {
	// Route: GET /users/{userID} → calls getUserHandler
	r.Get("/{userID}", getUserHandler)
}

// getUserHandler handles requests to /users/{userID}
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the "userID" path parameter from the URL
	userID := chi.URLParam(r, "userID")

	// Try to parse userID into an integer to validate it
	if _, err := strconv.Atoi(userID); err != nil {
		// If parsing fails, return HTTP 400 Bad Request
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// If valid, respond with user info
	fmt.Fprintf(w, "🧑 User Profile: %s\n", userID)
}

/*
🧠 LESSON 21 - ROUTING: USERS WITH DYNAMIC PATH PARAMETERS

✅ What You Learn:
- How to define dynamic route parameters using chi (e.g., /users/{userID})
- How to extract URL parameters cleanly inside a handler
- How to validate incoming path variables (basic integer validation)
- How to properly return client errors (HTTP 400) if validation fails

✅ Why This Matters:
- Dynamic URL parameters are foundational for real APIs (e.g., /users/{id}, /products/{sku})
- Path parsing and validation prevent bad input and improve server robustness
- Clean separation of domain logic makes it easier to test and extend handlers
- Professional Go servers expect validation and clean error handling at the route layer

Dynamic routing and clean parameter extraction like this prepare you for building APIs, microservices, and scalable backend systems.

*/
