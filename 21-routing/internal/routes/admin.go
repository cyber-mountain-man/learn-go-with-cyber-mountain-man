package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// AdminRoutes registers the /admin routes on the provided chi.Router
func AdminRoutes(r chi.Router) {
	// Route: GET /admin/dashboard → calls dashboardHandler
	r.Get("/dashboard", dashboardHandler)
}

// dashboardHandler handles requests to /admin/dashboard
func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	// Write a simple welcome message to the response
	fmt.Fprintln(w, "🛡️ Welcome to the Admin Dashboard!")
}

/*
🧠 LESSON 21 - ROUTING: ADMIN ROUTES

✅ What You Learn:
- How to group related paths (e.g., /admin/*) under a single function
- How to define clean RESTful route mappings with chi
- How to return basic HTML/text responses via handlers
- How chi's router chaining simplifies modular route registration

✅ Why This Matters:
- Grouping by domain (like "admin") prevents clutter in main.go
- Prepares you to extend /admin with more subroutes later (e.g., /admin/settings)
- Mirrors scalable API patterns used in real production systems
- Keeps route wiring highly readable and maintainable

This small file pattern becomes incredibly powerful once you add authentication, admin-only permissions, 
or grow your server with additional services.

*/
