package main

import (
    "log"              // For logging errors and server status
    "net/http"         // Provides HTTP server functionality

    "github.com/joho/godotenv" // Loads environment variables from .env file

    // Internal packages
    "github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/28-deployment/internal/db"
    "github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/28-deployment/internal/router"
)

func main() {
    // Load environment variables from .env (e.g., DB credentials)
    if err := godotenv.Load(); err != nil {
        log.Fatal("❌ Failed to load .env file") // Stop app if .env is missing
    }

    // Initialize the global database connection using environment config
    if err := db.InitDB(); err != nil {
        log.Fatalf("❌ Database connection failed: %v", err)
    }
    defer db.DB.Close() // 🔌 Ensure DB connection is closed on exit

    // Setup all application routes (static files, /users API, etc.)
    r := router.SetupRouter()

    // Start the server and listen on port 8080
    log.Println("🚀 Server running at http://localhost:8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatalf("❌ Server failed to start: %v", err)
    }
}

/*
Blurb: What This File Does
This is the entry point of your Go application. When the program starts, it follows a clear sequence:

Loads environment variables using the godotenv package. These values (like DBUSER, DBPASS, etc.) configure your database securely without hardcoding.

Initializes the global database connection through a reusable helper InitDB() in the db package. This allows the rest of your app to use the same open DB connection.

Sets up routing using chi, connecting URL endpoints to handler functions for things like serving static files and user management.

Starts the web server on localhost:8080. This is where users can interact with the application from their browser.

The combination of modular design (internal/db, internal/router, etc.) and clean startup flow makes this project easy to maintain and scalable for future features like authentication or API versioning.
*/