package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// PublicHandler handles requests to the /puplic route.
// No authentication is reuqires to access the endpoint.
func PublicHandler(w http.ResponseWriter, r *http.Request) {
	// Log the request method and path for obdervability
	log.Printf("[%s] %s", r.Method, r.URL.Path)

	// Send a basic public response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Welcome to the public endpoint!")
}

// PrivateHandler handles requests to the /private route.
// This is a protected endpoint that requires authentication via middleware.
func PrivateHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s", r.Method, r.URL.Path)

	// Simulate sensitive content that requires auth
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Welcome to the private endpoint! You are authenticated.")
}

/*
🧠 LESSON 23 - HANDLERS FOR PUBLIC & PROTECTED ROUTES

✅ What This Demonstrates:
- How to define basic HTTP handlers in Go
- The difference between unprotected and middleware-protected routes
- How to set response headers, status codes, and log structured access events

📌 Handler Best Practices:
- Keep handlers **lightweight** — focus on responding, not logic
- Use middleware to separate cross-cutting concerns like auth and logging
- Log all requests for visibility and auditability

🔐 Security Tip:
The `PrivateHandler` will only run if `RequireAuth` middleware passes —  
this reinforces **layered security**, a core principle in web backend design.

This file showcases clean handler separation and teaches how middleware and handlers
interact to build clear and maintainable web service flows.
*/
