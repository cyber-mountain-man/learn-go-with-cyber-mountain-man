package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// HomeHandler handles requests to the root "/" route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Log the HTTP method and URL path of the incoming request
	log.Printf("[%s] %s", r.Method, r.URL.Path)

	// If a subpath like "/something" hits this handler, redirect to a custom 404
	// This makes "/" behave like a strict match, not a prefix match
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}

	// Respond with HTTP 200 OK
	w.WriteHeader(http.StatusOK)

	// Write a welcome message to the response body
	fmt.Fprintln(w, "🏠 Welcome to the Home Page!")
}

// AboutHandler handles requests to the "/about" route
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	// Log request method and path
	log.Printf("[%s] %s", r.Method, r.URL.Path)

	// Return a successful status code
	w.WriteHeader(http.StatusOK)

	// Provide descriptive content about the server
	fmt.Fprintln(w, "This server is built with Go — clean, fast, and professional!")
}

// HealthHandler responds to "/health" requests for system monitoring
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	// Log health check access for observability tools
	log.Printf("[%s] %s", r.Method, r.URL.Path)

	// Return HTTP 200 OK to indicate the server is up
	w.WriteHeader(http.StatusOK)

	// Respond with a simple health message
	fmt.Fprintln(w, "✅ Server is healthy")
}

// NotFoundHandler handles undefined routes and returns a custom 404 page
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	// Log 404 errors with a special marker
	log.Printf("[%s] %s (404)", r.Method, r.URL.Path)

	// Return HTTP 404 Not Found
	w.WriteHeader(http.StatusNotFound)

	// Return custom error message
	fmt.Fprintln(w, "🚫 404 - Page Not Found")
}

/*
🧠 LESSON 20 - HANDLERS: ROUTES LOGIC

This file groups all HTTP route handling logic for the server.

✅ What You Learn:
- Each route ("/", "/about", "/health") gets its own clean handler function.
- Logging is used to track incoming HTTP methods and routes — critical for debugging and monitoring.
- Proper usage of HTTP status codes (200, 404) makes the app API-friendly and standards-compliant.
- Redirects for misused URLs show how Go routes are matched and handled with intent.
- The `/health` endpoint introduces real-world patterns for production monitoring systems (e.g., AWS, Kubernetes, GCP).

✅ Why This Matters:
- Cleanly separated handlers make the codebase easier to read, test, and extend.
- Encourages modular thinking: handlers are composable and can grow independently.
- Trains developers to treat logging, status codes, and route patterns seriously.
- Demonstrates how to begin abstracting route logic away from main.go.

These are foundational habits for backend Go engineers working on maintainable web APIs or services.
*/
