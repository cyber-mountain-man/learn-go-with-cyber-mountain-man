package main

import (
	"fmt"
	"log"
	"net/http"
)

// homeHandler handles requests to the "/" route
func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Home page accessed") // Log the request for monitoring
	fmt.Fprintln(w, "Welcome to the Go Web Server! 🖖") // Send a simple response
}

// aboutHandler handles requests to the "/about" route
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("About page accessed")
	fmt.Fprintln(w, "This server is built with Go. Clean, fast, and fun!")
}

// healthHandler handles requests to the "/health" route
func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Health check accessed")
	w.WriteHeader(http.StatusOK) // Set HTTP status 200 explicitly
	fmt.Fprintln(w, "OK")         // Standard health check response
}

func main() {
	// -----------------------------
	// 1️⃣ ROUTE SETUP
	// -----------------------------
	// Associate URL paths with their handler functions
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/health", healthHandler)

	// -----------------------------
	// 2️⃣ SERVER START
	// -----------------------------
	port := ":8080"
	log.Println("🚀 Server running at http://localhost" + port)

	// Start the HTTP server
	// - First argument: port to listen on
	// - Second argument: nil uses the default ServeMux (automatic route matching)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Server failed to start:", err) // Cleanly log and exit on fatal server errors
	}
}

/*
🧠 GO WEB SERVER — WALKTHROUGH

✅ What Happens Here:
- We define three handler functions, each responsible for a specific route ("/", "/about", "/health").
- Each handler responds to HTTP requests and logs activity to the console.
- We register these handlers to specific URL paths using `http.HandleFunc`.
- The server starts with `http.ListenAndServe` and listens for incoming requests on port 8080.

✅ Why This Matters:
- Building small web servers like this forms the **foundation** for APIs, web apps, microservices, and distributed systems in Go.
- `net/http` is Go's **standard library**, no external frameworks needed — powerful, minimal, and production-grade by default.
- Good logging and health check endpoints are **best practices** in real-world systems (monitoring, scaling, load balancing).

✅ Key Concepts:
| Concept                   | Explanation |
|----------------------------|-------------|
| `http.HandleFunc`          | Registers a function to handle HTTP requests on a given route |
| `http.ResponseWriter`      | Interface for sending responses to clients |
| `*http.Request`            | Contains all request info (method, URL, headers, etc.) |
| `log.Fatal` vs `fmt.Println` | `log.Fatal` exits the app on critical server errors |
| Health Check `/health`     | Real-world production servers expose `/health` endpoints for monitoring systems |

🔔 Bonus Tip:
- This server will **automatically handle concurrent connections** out of the box.
- Each incoming HTTP request is handled in its **own goroutine** — Go is built for concurrency, even at the web layer!

📚 Up Next:
- Add your own routes (e.g., `/contact`, `/services`)
- Serve static files (HTML, CSS)
- Build a simple REST API!

*/
