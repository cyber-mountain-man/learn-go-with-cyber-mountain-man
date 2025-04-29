package main

import (
	"log"
	"net/http"

	// Import our router setup logic
	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/22-controllers/internal/routes"
)

func main() {
	// -----------------------------
	// 1️⃣ SETUP ROUTER
	// -----------------------------
	// SetupRouter returns a fully configured chi.Router
	r := routes.SetupRouter()

	// -----------------------------
	// 2️⃣ START SERVER
	// -----------------------------
	port := ":8080"
	log.Println("🚀 Server running at http://localhost" + port)

	// Start HTTP server using the configured router
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal("Server failed:", err)
	}
}

/*
🧠 LESSON 22 - CONTROLLERS: SERVER ENTRY POINT

✅ What You Learn:
- How to bootstrap a Go web server using a chi.Router
- How to delegate route registration to a centralized router config
- How to start the server using clean separation of logic

✅ Why This Matters:
- Your main.go becomes clean, thin, and production-ready
- Keeps startup logic focused on orchestration, not handler logic
- Aligns with Go best practices for maintainability and scaling

This pattern lets your app grow into multiple services or modules without overwhelming main().
*/
