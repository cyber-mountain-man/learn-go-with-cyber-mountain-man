package main

import (
    "log"
    "net/http"

    "github.com/joho/godotenv"
    "github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/28-deployment/internal/db"
    "github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/28-deployment/internal/router"
)

func main() {
    // Load environment variables from .env
    if err := godotenv.Load(); err != nil {
        log.Fatal("❌ Failed to load .env file")
    }

    // Initialize global DB connection
    if err := db.InitDB(); err != nil {
        log.Fatalf("❌ Database connection failed: %v", err)
    }
    defer db.DB.Close()

    // Set up router
    r := router.SetupRouter() // formerly NewRouter(conn)

    log.Println("🚀 Server running at http://localhost:8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatalf("❌ Server failed to start: %v", err)
    }
}
