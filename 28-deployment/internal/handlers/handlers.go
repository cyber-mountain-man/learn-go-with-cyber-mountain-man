package handlers

import (
    "fmt"
    "net/http"
)

// Root handles the default route and confirms the app is working.
func Root(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("🌐 Deployed Go app with MSSQL is working!"))
}

// Health provides a simple health check endpoint.
func Health(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "✅ OK")
}