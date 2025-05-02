package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

// RenderHome handles HTTP GET requests to the home page
// and dynamically renders an HTML response using Go's html/template package.
func RenderHome(w http.ResponseWriter, r *http.Request) {
	// Log the request method and path (e.g., [GET] /)
	log.Printf("[%s] %s", r.Method, r.URL.Path)

	// -----------------------------
	// 1️⃣ TEMPLATE FILE PATHS
	// -----------------------------
	// Define the path to the base layout and page-specific content.
	layoutPath := filepath.Join("..", "..", "internal", "templates", "layout.html")
	contentPath := filepath.Join("..", "..", "internal", "templates", "home.html")

	// -----------------------------
	// 2️⃣ PARSE TEMPLATE FILES
	// -----------------------------
	// Combine layout + content into a single template object.
	tmpl, err := template.ParseFiles(layoutPath, contentPath)
	if err != nil {
		// If parsing fails (e.g., missing or malformed HTML), return a 500 error.
		http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// -----------------------------
	// 3️⃣ DYNAMIC DATA TO INJECT
	// -----------------------------
	// Create a map of data to be injected into the HTML template.
	data := map[string]interface{}{
		"Title": "Welcome to the Go Template Engine",
		"User":  "Mario",
		"Year":  time.Now().Year(), // <-- add this line
	}
	

	// -----------------------------
	// 4️⃣ EXECUTE & RENDER
	// -----------------------------
	// Execute the layout template (which includes home.html)
	// and write the final HTML to the response.
	if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
		http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
	}
}

/*
🧠 LESSON 24 - TEMPLATE RENDERING WALKTHROUGH

✅ What This Teaches:
- How to use Go’s built-in `html/template` engine for dynamic HTML rendering
- How to define reusable layout and content templates
- How to pass dynamic data into your views (e.g., map or struct)
- How to return complete, browser-ready HTML responses from handlers

🔍 How It Works:
- We use `template.ParseFiles()` to combine a base layout (e.g., nav, footer) with a content page
- Data is injected into placeholders using Go's templating syntax: {{ .Title }}, {{ .User }}
- Templates are safe by default — they auto-escape HTML to prevent XSS attacks

💡 Why It’s Useful:
- Allows your Go server to serve styled pages or views for dashboards, admin panels, etc.
- Keeps logic out of HTML and HTML out of Go — encourages separation of concerns
- Forms the foundation for server-rendered apps or admin tools built in pure Go

This is your first step toward building web UIs, portals, and full-featured Go-powered websites.
*/
