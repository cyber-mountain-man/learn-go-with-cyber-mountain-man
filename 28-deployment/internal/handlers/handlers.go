package handlers

import (
    "fmt"        // For formatted output (used in Health endpoint)
    "net/http"   // Go's HTTP server and client package
)

// Root handles the default route ("/").
// It sends a simple message confirming the app is deployed and responsive.
func Root(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("🌐 Deployed Go app with MSSQL is working!"))
}

// Health provides a minimal health check at "/health".
// It returns HTTP 200 OK with a confirmation message for uptime monitoring.
func Health(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)               // Explicitly set status to 200
    fmt.Fprintln(w, "✅ OK")                    // Send a success message to client
}

/*
🧠 Blurb: Understanding handlers/root.go
This file contains two basic HTTP handlers used for health checking and sanity testing:

Root(): Confirms that the application is running properly when someone accesses the root path (/). It’s useful as a quick visual test during deployment.

Health(): Used in automated systems (like Kubernetes probes or uptime bots) to verify that the server is live. It sends back a clean 200 OK status and message.

These routes are important for diagnostics and DevOps purposes but don’t perform any business logic. They help verify the container or server is up and responding to HTTP requests.
*/