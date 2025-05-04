package models

import (
	"database/sql"
	"time"
)

// User represents a row in the users table
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// GetAllUsers retrieves all users from the MSSQL database
func GetAllUsers(db *sql.DB) ([]User, error) {
	users := []User{}

	query := `SELECT id, name, email, created_at FROM users ORDER BY id`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, rows.Err()
}

// InsertUser inserts a new user and returns the new ID and CreatedAt timestamp
func InsertUser(db *sql.DB, u User) (int, error) {
	query := `
	INSERT INTO users (name, email)
	OUTPUT INSERTED.id, INSERTED.created_at
	VALUES (@p1, @p2)`

	var id int
	var createdAt time.Time

	err := db.QueryRow(query, u.Name, u.Email).Scan(&id, &createdAt)
	u.CreatedAt = createdAt

	return id, err
}

// GetUserByID fetches a user by their ID
func GetUserByID(db *sql.DB, id int) (User, error) {
	var u User

	query := `SELECT id, name, email, created_at FROM users WHERE id = @p1`
	err := db.QueryRow(query, id).Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)

	return u, err
}

// UpdateUser modifies the name and email for a given user
func UpdateUser(db *sql.DB, u User) error {
	query := `UPDATE users SET name = @p1, email = @p2 WHERE id = @p3`
	_, err := db.Exec(query, u.Name, u.Email, u.ID)
	return err
}

// DeleteUser removes a user by ID
func DeleteUser(db *sql.DB, id int) error {
	query := `DELETE FROM users WHERE id = @p1`
	_, err := db.Exec(query, id)
	return err
}
