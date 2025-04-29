package main

import (
	"log"
	"net/http"

	// Import handler functions from internal package
	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/20-handlers/internal/handlers"
)

func main() {
	// -----------------------------
	// 1️⃣ EXPLICIT ROUTER SETUP
	// -----------------------------
	// Use a custom ServeMux instead of the global DefaultServeMux.
	// This improves modularity and avoids hidden global state.
	mux := http.NewServeMux()

	// -----------------------------
	// 2️⃣ ROUTE REGISTRATION
	// -----------------------------
	// Define paths and associate them with handler functions.
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/about", handlers.AboutHandler)
	mux.HandleFunc("/health", handlers.HealthHandler) // Health check for uptime monitoring
	mux.HandleFunc("/404", handlers.NotFoundHandler)  // Fallback route

	// -----------------------------
	// 3️⃣ SERVER STARTUP
	// -----------------------------
	port := ":8080"
	log.Println("🚀 Server running at http://localhost" + port)

	// Start the server using the custom router (mux)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

/*
🧠 LESSON 20 - HANDLERS: MAIN ENTRY POINT

This file demonstrates a clean, modular way to launch a Go web server:

✅ What You Learn:
- How to replace http.DefaultServeMux with your own ServeMux for better testability
- How to wire up route handlers without mixing them into the main entry file
- How to create a clean and thin main.go that only initializes routing and the server
- How to add a standard /health endpoint for monitoring readiness/liveness
- Proper use of logging and graceful startup error handling

✅ Why This Matters:
- This layout mirrors what professional Go engineers use in production apps
- It's easier to maintain and scale — no business logic clutters main.go
- It's ready for future enhancements like graceful shutdown, middleware, or dependency injection

This file is the first step to building highly structured and reliable web applications in Go.
*/
