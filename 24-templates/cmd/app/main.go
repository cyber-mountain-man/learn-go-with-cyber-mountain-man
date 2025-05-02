package main

import (
	"log"
	"net/http"

	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/24-templates/internal/handlers"
)

func main() {
	// -----------------------------
	// 1️⃣ ROUTE SETUP
	// -----------------------------
	// Register route paths with their corresponding handler functions.
	// These handlers are defined in the internal/handlers package.
	// Each route maps a URL to a function that writes an HTTP response.

	http.HandleFunc("/", handlers.RenderHome)       // Serves the HTML template
	http.HandleFunc("/public", handlers.PublicHandler) // Public access message
	http.HandleFunc("/private", handlers.PrivateHandler) // Simulated protected route

	// -----------------------------
	// 2️⃣ SERVER CONFIGURATION
	// -----------------------------
	// Define the port the server should listen on.
	port := ":8080"

	// Log to the terminal where the server is running.
	log.Println("🌐 Template server running at http://localhost" + port)

	// -----------------------------
	// 3️⃣ START THE SERVER
	// -----------------------------
	// Start listening for incoming HTTP requests on the specified port.
	// If an error occurs (e.g., port already in use), log and exit.
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("❌ Server failed to start:", err)
	}
}

/*
🧠 LESSON 24 - GO HTTP SERVER ENTRYPOINT

✅ What You Learn:
- How to set up HTTP routes in Go using `http.HandleFunc`
- How to register multiple endpoints (`/`, `/public`, `/private`)
- How to start and monitor a basic server with logging

📌 Real-World Context:
This entry point acts as the central control point of a Go web app.
You define the server configuration, map request paths, and initialize runtime behavior.

🔍 Key Takeaways:
- Use `http.HandleFunc` for lightweight routing
- Log the startup port to track server availability
- Modularize business logic into external handler packages

🚀 What's Next:
- Replace `http.HandleFunc` with `chi.Router` or `net/http.ServeMux` for more advanced routing
- Add middleware to secure the `/private` route
- Dynamically serve multiple pages with reusable layouts and templates

This main function is the heart of your server — it boots up your application and binds it to real-world traffic.
*/
