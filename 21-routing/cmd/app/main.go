package main

import (
	"log"
	"net/http"

	// Import our internal route groups
	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/21-routing/internal/routes"

	// Import chi router
	"github.com/go-chi/chi/v5"
)

func main() {
	// -----------------------------
	// 1️⃣ INITIALIZE ROUTER
	// -----------------------------
	// chi.NewRouter creates a fresh instance of a router.
	// It's lightweight and provides flexible middleware support.
	r := chi.NewRouter()

	// -----------------------------
	// 2️⃣ MOUNT ROUTE GROUPS
	// -----------------------------
	// Attach specific route handlers grouped by domain logic
	r.Route("/users", routes.UserRoutes)  // Handles /users/{userID}
	r.Route("/admin", routes.AdminRoutes) // Handles /admin/dashboard

	// -----------------------------
	// 3️⃣ HANDLE 404 NOT FOUND
	// -----------------------------
	// Customize behavior for unknown routes
	r.NotFound(routes.FallbackHandler)

	// -----------------------------
	// 4️⃣ START SERVER
	// -----------------------------
	port := ":8080"
	log.Println("🚀 Server running at http://localhost" + port)

	// Start the server with the chi router
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal("Server failed:", err)
	}
}

/*
🧠 LESSON 21 - ROUTING: MAIN SERVER SETUP

✅ What You Learn:
- How to initialize a modern, flexible router using chi
- How to group routes cleanly by domain (e.g., /users, /admin)
- How to set up a fallback 404 handler for better UX
- How to start a professional Go web server modularly

✅ Why This Structure Matters:
- Route grouping keeps server logic clean and scalable
- Using chi provides a lightweight, battle-tested routing system
- Separating the server wiring (main.go) from route logic (internal/routes) mirrors real production systems
- Makes your server easier to extend, test, and deploy

This server structure sets the foundation for real-world Go APIs and web services.
It prepares you for adding features like authentication, middleware, and even versioned APIs (v1, v2).

*/
