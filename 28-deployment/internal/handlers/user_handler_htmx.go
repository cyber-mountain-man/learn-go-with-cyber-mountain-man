package handlers

import (
	"net/http"                              // Standard HTTP utilities
	"strconv"                               // For string-to-int conversion (e.g., ID parsing)
	"html/template"                         // HTML templating for rendering fragments
	"log"                                   // Logging for debug and error output
	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/28-deployment/internal/db"
	"github.com/go-chi/chi/v5"             // Router for extracting path parameters like /users/{id}
)

// Precompile all HTML templates in the static/templates directory
var templates = template.Must(template.ParseGlob("static/templates/*.html"))

// ListUsersHTMX renders all users using the user-list.html fragment.
// This is triggered via HTMX GET and used to refresh the full user list.
func ListUsersHTMX(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	templates.ExecuteTemplate(w, "user-list.html", users)
}

// CreateUserHTMX handles the HTMX form submission for adding a user.
// If successful, it returns the refreshed user list.
func CreateUserHTMX(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form submission", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	email := r.FormValue("email")

	if err := db.InsertUser(name, email); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Return updated list to HTMX
	ListUsersHTMX(w, r)
}

// EditUserFormHTMX returns an edit form populated with user data.
// HTMX injects this into the DOM dynamically.
func EditUserFormHTMX(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := db.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Inject the edit form fragment into the page
	err = templates.ExecuteTemplate(w, "user-edit.html", user)
	if err != nil {
		log.Println("Template execution error:", err)
		http.Error(w, "Template execution failed: "+err.Error(), http.StatusInternalServerError)
	}
}

// UpdateUserHTMX processes the PUT request from the edit form and updates the user.
// After updating, it refreshes the user list (used in HTMX swap).
func UpdateUserHTMX(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	email := r.FormValue("email")

	if err := db.UpdateUser(id, name, email); err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	// Re-render user list after successful update
	ListUsersHTMX(w, r)
}

// DeleteUserHTMX removes a user from the database based on the URL param ID.
// The result is a refreshed list returned to HTMX.
func DeleteUserHTMX(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := db.DeleteUser(id); err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	ListUsersHTMX(w, r)
}


/*
🧠 Blurb: Understanding user_handler_htmx.go
This file defines all HTMX-compatible CRUD HTTP handlers for the /users resource in your Go web app:

HTMX Role: HTMX triggers server calls using GET, POST, PUT, or DELETE, and expects HTML snippets (not full pages) in return. This enables partial updates to the DOM.

Server Role: Each handler performs DB operations and then renders the appropriate HTML fragment using Go templates.

Templates Used:

user-list.html: Displays all users.

user-edit.html: Editable form injected during the Edit cycle.

This design provides a lightweight frontend-backend interaction without needing a SPA framework like React or Vue. It's a powerful, simple pattern for dynamic UIs.
*/