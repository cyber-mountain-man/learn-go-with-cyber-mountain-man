package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"learn-go-with-cyber-mountain-man/26-databases/mssql/internal/models"
)

// ListUsers handles GET /users and returns all users
func ListUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := models.GetAllUsers(db)
		if err != nil {
			http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(users)
	}
}

// CreateUser handles POST /users and inserts a new user
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

// GetUser handles GET /users/{id} and returns a user by ID
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
			http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(u)
	}
}

// UpdateUser handles PUT /users/{id} to update an existing user
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

		json.NewEncoder(w).Encode(u)
	}
}

// DeleteUser handles DELETE /users/{id}
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
