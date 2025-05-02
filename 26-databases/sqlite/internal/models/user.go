package models

import (
	"database/sql"
)

// User represents a user in the system.
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUser inserts a new user into the database.
func CreateUser(db *sql.DB, user *User) error {
	query := `INSERT INTO users (name, email) VALUES (?, ?)`
	result, err := db.Exec(query, user.Name, user.Email)
	if err != nil {
		return err
	}
	lastID, err := result.LastInsertId()
	if err == nil {
		user.ID = int(lastID)
	}
	return nil
}

// UpdateUser modifies an existing user record.
func UpdateUser(db *sql.DB, user *User) error {
	query := `UPDATE users SET name = ?, email = ? WHERE id = ?`
	_, err := db.Exec(query, user.Name, user.Email, user.ID)
	return err
}

// DeleteUser removes a user from the database by ID.
func DeleteUser(db *sql.DB, id int) error {
	_, err := db.Exec(`DELETE FROM users WHERE id = ?`, id)
	return err
}

// GetAllUsers fetches all users from the database.
func GetAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query(`SELECT id, name, email FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// GetUserByID retrieves a user by their ID.
func GetUserByID(db *sql.DB, id int) (*User, error) {
	var u User
	query := `SELECT id, name, email FROM users WHERE id = ?`
	err := db.QueryRow(query, id).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
