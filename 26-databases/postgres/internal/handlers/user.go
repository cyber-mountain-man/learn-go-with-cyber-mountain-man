package handlers


import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"learn-go-with-cyber-mountain-man/26-databases/postgres/internal/models"
)

// ListUsers returns all users
func ListUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := models.GetAllUsers(db)
		if err != nil {
			http.Error(w, "Error fetching users", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(users)
	}
}

// CreateUser adds a new user
func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u models.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		id, err := models.InsertUser(db, u)
		if err != nil {
			http.Error(w, "Failed to insert user", http.StatusInternalServerError)
			return
		}
		u.ID = id

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(u)
	}
}

// GetUser returns a user by ID
func GetUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		u, err := models.GetUserByID(db, id)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}
			http.Error(w, "Error retrieving user", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(u)
	}
}

// UpdateUser modifies an existing user
func UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		var u models.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		u.ID = id

		if err := models.UpdateUser(db, u); err != nil {
			http.Error(w, "Failed to update user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(u)
	}
}

// DeleteUser removes a user by ID
func DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		if err := models.DeleteUser(db, id); err != nil {
			http.Error(w, "Failed to delete user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

/*
🧠 GO + POSTGRESQL CRUD — ROUTE HANDLERS WALKTHROUGH (handlers/user.go)

✅ What Happens Here:
- We define HTTP handlers for all CRUD operations on the `/users` endpoint.
- These handlers delegate data operations to the `models` package and focus purely on request/response logic.
- Input from the client is decoded from JSON, and responses are returned as JSON using `encoding/json`.
- Route parameters (like `/users/{id}`) are extracted using `chi.URLParam`.

✅ Why This Matters:
- This file forms the **controller layer** in a clean architecture: it receives requests, calls model functions, and sends responses.
- You learn to use `http.HandlerFunc`, `chi.Router`, and `json.NewDecoder/Encoder` to build production-ready REST APIs.
- Handling errors properly and returning appropriate HTTP status codes (e.g., 200, 201, 400, 404, 500) is critical for API reliability.

✅ Key Concepts:
| Concept                    | Explanation |
|----------------------------|-------------|
| `http.HandlerFunc`         | Converts functions into valid HTTP handlers |
| `chi.URLParam`             | Extracts dynamic values from route paths |
| `json.NewDecoder`          | Parses JSON from request bodies into Go structs |
| `http.Error`               | Sends structured error responses with status codes |
| Status codes: 201, 404     | Signify created, not found, etc. — important for clients and debugging |

🔔 Bonus Tip:
- These handlers are stateless and pure — ideal for microservices or integration into larger systems.
- You can easily add middleware (e.g., logging, auth) later thanks to Chi’s composable design.

📚 Up Next:
- Add validation for input fields (e.g., check if name/email are empty)
- Extract common logic (like ID parsing) into helper functions
*/
