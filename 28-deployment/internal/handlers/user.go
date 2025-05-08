package handlers

import (
	"database/sql"                         // SQL database interactions
	"encoding/json"                        // JSON encoding/decoding
	"fmt"                                  // String formatting for HTML output
	"net/http"                             // Core HTTP functionality
	"strconv"                              // Convert path variables (ID) to integers

	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/28-deployment/internal/models"
	"github.com/go-chi/chi/v5"            // Router library for path parameters
)

// UserHandler groups database access so it can be injected and reused
type UserHandler struct {
	DB *sql.DB
}

// GetAllUsers serves both:
// - JSON (for API calls)
// - HTML (if `Accept: text/html` or `HX-Request` is present)
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	accept := r.Header.Get("Accept")
	if accept == "text/html" || r.Header.Get("HX-Request") == "true" {
		h.RenderUserList(w, r)
		return
	}

	// JSON fallback: used in pure REST scenarios
	rows, err := h.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		http.Error(w, "Database query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			http.Error(w, "Failed to scan row", http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUser returns one user as JSON by ID
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var u models.User
	err = h.DB.QueryRow("SELECT id, name, email FROM users WHERE id = @p1", id).
		Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(u)
}

// CreateUser handles both:
// - HTMX form submissions (via r.FormValue)
// - JSON API posts (via json.NewDecoder)
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User

	// Try JSON first
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil || u.Name == "" || u.Email == "" {
		// Fallback to form values (common for HTMX)
		u.Name = r.FormValue("name")
		u.Email = r.FormValue("email")
	}

	if u.Name == "" || u.Email == "" {
		http.Error(w, "Name and email are required fields", http.StatusBadRequest)
		return
	}

	// Insert user and return new ID
	err = h.DB.QueryRow(
		"INSERT INTO users (name, email) OUTPUT INSERTED.id VALUES (@p1, @p2)",
		u.Name, u.Email,
	).Scan(&u.ID)
	if err != nil {
		http.Error(w, "Insert failed", http.StatusInternalServerError)
		return
	}

	// Refresh the user list for HTMX
	h.RenderUserList(w, r)
}

// UpdateUser allows users to be updated via JSON
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	if u.Name == "" || u.Email == "" {
		http.Error(w, "Name and email are required fields", http.StatusBadRequest)
		return
	}

	_, err := h.DB.Exec(
		"UPDATE users SET name = @p1, email = @p2 WHERE id = @p3",
		u.Name, u.Email, id,
	)
	if err != nil {
		http.Error(w, "Update failed", http.StatusInternalServerError)
		return
	}

	// No HTML returned; used in API or async HTMX context
	w.WriteHeader(http.StatusNoContent)
}

// DeleteUser removes the user and refreshes the list for HTMX
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	_, err = h.DB.Exec("DELETE FROM users WHERE id = @p1", id)
	if err != nil {
		http.Error(w, "Delete failed", http.StatusInternalServerError)
		return
	}

	// Refresh the list after deletion
	h.RenderUserList(w, r)
}

// RenderUserList creates an HTML unordered list of users (for HTMX swap)
func (h *UserHandler) RenderUserList(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		http.Error(w, "Failed to load users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, `<ul id="user-list">`)
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err == nil {
			fmt.Fprintf(
				w,
				`<li>%s (%s) <button class="delete-btn" hx-delete="/users/%d" hx-target="#user-list" hx-swap="outerHTML">Delete</button></li>`,
				u.Name, u.Email, u.ID,
			)
		}
	}
	fmt.Fprintln(w, `</ul>`)
}

/*
🧠 Blurb: Understanding UserHandler (API + HTMX Hybrid)
This handler serves two purposes:

RESTful API: Supports JSON input/output via standard Accept headers and json.NewDecoder.

HTMX Support: Dynamically serves HTML snippets like <ul> and handles form-based interactions.

Key Capabilities:

Dual handling of form data and JSON input.

Dynamically updated frontend via HTMX (hx-get, hx-post, hx-delete, etc.).

Automatic user list refreshing after changes (Create/Delete).

Graceful fallback to JSON output when HTMX is not used.

You can use this in parallel with a full HTMX interface and REST API consumers (e.g., Postman, frontend apps), making it flexible and powerful.
*/