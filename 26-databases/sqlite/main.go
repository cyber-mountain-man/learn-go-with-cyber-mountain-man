package main

import (
	"log"
	"net/http"

	"sqlite/internal/db"
	"sqlite/internal/routes"
)

func main() {
	// -----------------------------
	// 1️⃣ INITIALIZE DATABASE
	// -----------------------------
	if err := db.InitDB("data/app.db"); err != nil {
		log.Fatal("❌ Failed to initialize database:", err)
	}

	// -----------------------------
	// 2️⃣ SET UP ROUTES USING CHI
	// -----------------------------
	router := routes.Register(db.DB) // Inject db into router setup

	// -----------------------------
	// 3️⃣ START SERVER
	// -----------------------------
	port := ":8080"
	log.Println("🚀 Server running at http://localhost" + port)

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal("❌ Server failed to start:", err)
	}
}