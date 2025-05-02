package handlers

import (
	"fmt"
	"net/http"
)

// PublicHandler handles GET requests to the /public route.
// This route simulates a page that is accessible to all users,
// without requiring authentication or special access.
func PublicHandler(w http.ResponseWriter, r *http.Request) {
	// Write a simple response to the browser or client
	fmt.Fprintln(w, "🌐 This is a public page. Everyone can see it.")
}

// PrivateHandler handles GET requests to the /private route.
// This simulates a restricted resource — in a real app,
// this would be protected by middleware or login validation.
func PrivateHandler(w http.ResponseWriter, r *http.Request) {
	// Return a message indicating access is restricted
	fmt.Fprintln(w, "🔐 This is a protected page. You must be authenticated to view this.")
}

/*
🧠 LESSON 24 EXTENSION – PUBLIC & PRIVATE ROUTES

✅ What You Learn:
- How to create simple route handlers using the net/http package
- The difference between unrestricted (public) and restricted (private) content
- How to return plain text responses to HTTP clients

💡 Real-World Usage:
- Public routes are useful for marketing pages, help sections, or open APIs
- Private routes typically require middleware for session validation, tokens, or user roles
- This sets the stage for building secured applications using Go

📌 Next Steps:
- Enhance /private with basic authentication (e.g., check headers)
- Apply custom middleware (like `RequireAuth`) to simulate protected access
*/