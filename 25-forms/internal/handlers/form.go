package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// RenderForm displays the form to the user via GET /form
func RenderForm(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s", r.Method, r.URL.Path)

	// Load layout and form templates
	layout := filepath.Join("internal", "templates", "layout.html")
	form := filepath.Join("internal", "templates", "form.html")

	tmpl, err := template.ParseFiles(layout, form)
	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Optional: Pass dynamic page data
	data := map[string]interface{}{
		"Title": "User Feedback Form",
	}

	if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
		http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
	}
}

// HandleForm processes form submission via POST /submit
func HandleForm(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s", r.Method, r.URL.Path)

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/form", http.StatusSeeOther)
		return
	}

	// Parse form values from POST body
	name := r.FormValue("name")
	email := r.FormValue("email")
	message := r.FormValue("message")

	// Basic validation (could be expanded)
	if name == "" || email == "" || message == "" {
		http.Error(w, "All fields are required.", http.StatusBadRequest)
		return
	}

	// Output confirmation (normally you'd store it or email it)
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h2>✅ Form submitted successfully!</h2>"))
	w.Write([]byte("<p><strong>Name:</strong> " + name + "</p>"))
	w.Write([]byte("<p><strong>Email:</strong> " + email + "</p>"))
	w.Write([]byte("<p><strong>Message:</strong> " + message + "</p>"))
}

/*
🧠 LESSON 25 - FORM HANDLER LOGIC

✅ What You Learn:
- How to render an HTML form using Go templates
- How to extract POST data with `r.FormValue()`
- How to perform simple input validation
- How to render a confirmation response

📦 Real-World Use Cases:
- Contact forms
- Admin data entry
- Survey tools, feedback modules

🛡 Best Practices:
- Validate required fields server-side
- Escape dynamic output when embedding into templates (for security)
- Redirect GET requests to POST-only routes

This pattern mirrors how you'd process and confirm form input in Go-based production servers.
*/
