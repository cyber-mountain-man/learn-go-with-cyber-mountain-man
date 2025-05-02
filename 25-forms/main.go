package main

import (
	"log"
	"net/http"

	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/25-forms/internal/handlers"
)

func main() {
	// -----------------------------
	// 1️⃣ ROUTE SETUP
	// -----------------------------

	// Redirect root (/) to /form for cleaner UX
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/form", http.StatusSeeOther)
	})

	// GET /form → Show the form
	http.HandleFunc("/form", handlers.RenderForm)

	// POST /submit → Handle form submission
	http.HandleFunc("/submit", handlers.HandleForm)

	// -----------------------------
	// 2️⃣ SERVER CONFIGURATION
	// -----------------------------
	port := ":8080"
	log.Println("📬 Form server running at http://localhost" + port)

	// -----------------------------
	// 3️⃣ START THE SERVER
	// -----------------------------
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("❌ Server failed to start:", err)
	}
}

/*
🧠 LESSON 25 - HTML FORMS IN GO

✅ What This Teaches:
- How to serve an HTML form using `http.HandleFunc`
- How to handle form submission with POST routes
- How to separate GET and POST logic cleanly
- How to redirect `/` to a defined route (`/form`)

🔍 Real-World Relevance:
This mirrors common patterns in dashboards, admin panels, and CMS tools.
Go's standard library makes it easy to build secure, high-performance form handling.

📌 Best Practices:
- Always validate and sanitize form data server-side
- Redirect root paths to logical start points (like `/form`)
- Keep logic clean by splitting GET/POST responsibilities
- Use `template.ParseFiles()` to render form pages dynamically

This is a foundational pattern in Go web apps — clean, predictable, and built for scale.
*/
