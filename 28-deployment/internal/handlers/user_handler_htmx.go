package handlers

import (
	"net/http"
	"strconv"
	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/28-deployment/internal/db"
	"html/template"
	"github.com/go-chi/chi/v5"
)

var templates = template.Must(template.ParseGlob("static/templates/*.html"))

func ListUsersHTMX(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	templates.ExecuteTemplate(w, "user-list.html", users)
}

func CreateUserHTMX(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form submission", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	email := r.FormValue("email")

	err = db.InsertUser(name, email)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// ✅ Return updated user list
	ListUsersHTMX(w, r)
}

func EditUserFormHTMX(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)
	user, err := db.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	templates.ExecuteTemplate(w, "user-edit.html", user)
}

func UpdateUserHTMX(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	email := r.FormValue("email")

	err = db.UpdateUser(id, name, email)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}
	ListUsersHTMX(w, r)
}

func DeleteUserHTMX(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	err := db.DeleteUser(id)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}
	ListUsersHTMX(w, r)
}
