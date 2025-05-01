package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logger is a global middleware that logs each incoming HTTP request.
// It wraps the next handler in the chain, allowing logging before and after the route logic.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Record the time when the request starts (for performance metrics)
		start := time.Now()

		// Log the request method and requested URL path
		// Example output: [GET] /public
		log.Printf("[%s] %s", r.Method, r.URL.Path)

		// Pass control to the next handler (this is required)
		// Without this, the request would be blocked here.
		next.ServeHTTP(w, r)

		// After the handler finishes, log how long the request took to process
		// This helps track slow requests and overall performance
		log.Printf("↪ Completed in %s\n", time.Since(start))
	})
}

/*
🧠 LESSON 23 - LOGGER MIDDLEWARE WALKTHROUGH

🔍 What This Middleware Does:
The `Logger` function is a standard HTTP middleware in Go.
It wraps every incoming request and:

1. Captures the **request method and path** — e.g., [GET] /public
2. Measures how **long it takes** for the entire handler chain to respond
3. Outputs structured log lines before and after each request

💡 How It Works:
- `Logger` returns a function that conforms to the `http.Handler` interface
- It uses `next.ServeHTTP(w, r)` to ensure the next handler or middleware is called
- You can stack it with other middleware (like authentication, rate limiting, etc.)

🔐 Why This Is Useful in Real Projects:
- Helps monitor traffic and behavior in production (e.g., identify slow endpoints)
- Acts as a **non-intrusive debug tool** — no need to modify handlers to trace requests
- Can be extended to include more info like IP address, status codes, headers, etc.

📦 Common Use Cases:
- Add to every API in production to support observability
- Pair with logging platforms like Loki, Datadog, or AWS CloudWatch
- Use timing data to detect latency spikes and regressions

This pattern is the backbone of middleware composition in Go —
every robust server you build will rely on wrappers like this.
*/