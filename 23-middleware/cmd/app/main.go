package main

import (
	"log"
	"net/http"

	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/23-middleware/internal/handlers"
	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/23-middleware/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	// -----------------------------
	// 1️⃣ CREATE ROUTER
	// -----------------------------
	// chi.NewRouter returns a fresh, composable router instance.
	// Think of this like the main "mux" that all requests pass through.
	r := chi.NewRouter()

	// -----------------------------
	// 2️⃣ GLOBAL MIDDLEWARE
	// -----------------------------
	// r.Use(middleware...) attaches middleware that will run for every route.
	// Logger will print method/path for all incoming requests.
	r.Use(middleware.Logger)

	// -----------------------------
	// 3️⃣ PUBLIC ROUTES (NO AUTH REQUIRED)
	// -----------------------------
	r.Group(func(r chi.Router) {
		// This group uses NO special middleware — just responds directly.
		r.Get("/public", handlers.PublicHandler)
	})

	// -----------------------------
	// 4️⃣ PROTECTED ROUTES (AUTH REQUIRED)
	// -----------------------------
	r.Group(func(r chi.Router) {
		// This middleware only applies to this block.
		// You can stack more here: e.g., r.Use(Throttle, CORS, etc.)
		r.Use(middleware.RequireAuth)

		// Simulated protected route
		r.Get("/private", handlers.PrivateHandler)
	})

	// -----------------------------
	// 5️⃣ START SERVER
	// -----------------------------
	port := ":8080"
	log.Println("🚀 Server running at http://localhost" + port)

	// Start the HTTP server using our composed router
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal("❌ Server failed to start:", err)
	}
}

/*
🧠 LESSON 23 - MIDDLEWARE ENTRYPOINT

✅ What You Learn:
- How to declare middleware using `r.Use()` in chi
- How to group routes with different middleware stacks
- How to cleanly separate concerns: server, routes, logic, and middleware

✅ Why This Matters:
- Middleware is how you log, secure, or modify requests in production systems
- This design allows selective use of logic (e.g. logging, auth) without duplication
- Keeps your server flexible and maintainable
- Middleware powers core backend features: logging, auth, rate limiting, compression, etc.
- Go’s chi router supports *middleware layering*, which allows you to isolate sensitive logic only where needed.
- Makes your `main.go` clean and orchestration-focused (no logic, just routing + control).

This file defines the "control panel" for your server’s middleware strategy.
It’s clean, composable, and idiomatic — the Go way.

This structure is used in real SaaS products, APIs, and microservices.  
You’ll use this design when building secure, observable Go backends that scale cleanly.
*/
